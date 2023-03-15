package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
)

func lastChgTime(pathname string) int64 {
	finfo, _ := os.Stat(pathname)
	winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
	return winFileAttr.LastWriteTime.Nanoseconds() / 1e9
}

func execProtoc(pathname string) {
	path := CutPath(pathname)
	name := CutName(pathname)
	cmd := fmt.Sprintf("protoc --go_out=%s %s --proto_path %s", path, pathname, path)
	if !execCommand(cmd) {
		return
	}

	cppPath := path + "cpp/"
	os.Mkdir(cppPath, 0666)
	cmd = fmt.Sprintf("protoc --cpp_out=%s %s --proto_path %s", cppPath, pathname, path)
	if !execCommand(cmd) {
		return
	}

	pbgo := path + strings.Replace(name, ".proto", ".pb.go", -1)
	f, err := os.Open(pbgo)
	if err != nil {
		fmt.Println("文件打开错误:", pbgo)
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	filetxt := `
import (
	"lib/amc"
)

func init() {`

	rawText := ""
	pack := ""
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		rawText = rawText + line

		if strings.Index(line, "package ") == 0 {
			pack = line
		}

		const key1 = ".RegisterType("

		index := strings.Index(line, key1)
		if index == -1 {
			continue
		}

		s := line[index+len(key1):]

		const key2 = ", \""
		index = strings.Index(s, key2)
		if index == -1 {
			continue
		}
		filetxt = filetxt + "\n\tamc.RegisterProtoMsg(" + s[0:index] + ")"
	}

	filetxt = filetxt + `
}
`

	rawText = strings.Replace(rawText, "fileDescriptor0", "fileDescriptor_"+strings.Replace(name, ".proto", "", -1), -1)
	StrToFile(rawText, pbgo)

	if pack != "" {
		filetxt = pack + filetxt
	} else {
		filetxt = "package proto\n" + filetxt
	}
	StrToFile(filetxt, path+strings.Replace(name, ".proto", ".amc.go", -1))
	return
}

var lastChgTimes = make(map[string]int64)

func goWork(workPath string, loop bool) {
	for {
		files, err := WalkDir(workPath, "proto")
		if err != nil {
			fmt.Println("获取Proto文件失败：", err)
		}

		hasChg := false

		for _, file := range files {
			f := strings.ToLower(file)

			lastTime := lastChgTimes[f]
			newTime := lastChgTime(file)
			lastChgTimes[f] = newTime

			if newTime == lastTime {
				continue
			}

			hasChg = true
			fmt.Println("Proto文件更新：", file)
			execProtoc(file)
		}

		if !loop {
			break
		}

		if hasChg {
			OutCurPath()
		}

		time.Sleep(time.Duration(time.Second * 2))
	}
}

func doWork(workPath string) {
	goWork(workPath, false)
	go goWork(workPath, true)
}
