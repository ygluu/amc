package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

func getIniValue(filename, section, key string) string {
	if filename == "" {
		return ""
	}
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	key = strings.ToLower(key)
	section = strings.ToLower(section)
	reader := bufio.NewReader(file)
	var sname string = ""
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		if linestr[0] == ';' {
			continue
		}
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sname = linestr[1 : len(linestr)-1]
		} else if strings.ToLower(sname) == section {
			pair := strings.Split(linestr, "=")
			if len(pair) == 2 {
				kname := strings.TrimSpace(pair[0])
				if strings.ToLower(kname) == key {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

var exePath = CutPath(ExecPath())

func execCommand(scmd string) bool {
	fmt.Println(scmd)
	args := strings.Split(scmd, " ")
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("\nprotoc error:", err)
		fmt.Println(stderr.String())
		return false
	}
	if len(out.String()) > 0 {
		fmt.Println("\nproto out:\n" + out.String())
	}
	return true
}

func saveLastWorkDir(workdir string) {
	s := `[Config]
LastWorkPaht=%s
`
	workdir = strings.Replace(workdir, "/src", "", -1)
	StrToFile(fmt.Sprintf(s, workdir), exePath+"cfg.ini")
}

func getLastWorkDir() string {
	return getIniValue(exePath+"cfg.ini", "Config", "LastWorkPaht")
}

var workPath = getLastWorkDir()

func selWorkPath() {
	input := bufio.NewScanner(os.Stdin)

	workPath := getLastWorkDir()
	if workPath != "" {
		doWork(workPath)
		return
	}

	for {
		fmt.Println("请输入工作路径:")

		input.Scan()
		text := input.Text()
		if !PathExists(text) {
			fmt.Println("无效路径：", workPath)
			continue
		}
		workPath = text
		workPath = strings.Replace(workPath, "\\", "/", -1)
		saveLastWorkDir(workPath)
	}
	doWork(workPath)
}

func ExecCmd(cmd string) {
	fmt.Println(cmd)
}

func ParseIn(in string) (ret []string) {
	ss := strings.Split(in, " ")
	if len(ss) == 0 {
		return
	}

	for i := 0; i < len(ss)-1; i++ {
		s := ss[i]
		if ss[i] == "" {
			continue
		}
		ret = append(ret, s)
	}
	last := ss[len(ss)-1]
	ret = append(ret, last[0:len(last)-2])
	return
}

var curPath = workPath

func OutCurPath() {
	fmt.Print("[" + curPath + "]:")
}

func OutWorkPath() {
	fmt.Println("工作路径:", workPath)
}

type cmdInfo struct {
	handler func([]string)
	desc    string
}

var cmdHandlers = make(map[string]*cmdInfo)

func cmdDir(params []string) {
	dirs := DirList(curPath)
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func cmdCd(params []string) {
	if len(params) != 1 {
		fmt.Println("错误CD参数")
		return
	}

	dir := params[0]
	if dir == ".." {
		curPath = ToParentDir(curPath)
		return
	}

	path := curPath + dir + "/"

	if !PathExists(path) {
		fmt.Println("没有这个目录:", dir)
		return
	}

	curPath = path
}

func cmdWorkDir(params []string) {
	workPath = curPath
	saveLastWorkDir(workPath)
	fmt.Println("当前工作路径:", workPath)
}

func printHelp(params []string) {
	fmt.Println("命令说明：")
	fmt.Println("****************************************")
	fmt.Println("D:/xxx	:直接输入全路径名为当前目录，并默认为工作目录")

	list := []string{}
	for cmd, info := range cmdHandlers {
		list = append(list, fmt.Sprintf("%s	:%s", cmd, info.desc))
	}

	sort.Strings(list)
	for _, s := range list {
		fmt.Println(s)
	}
	fmt.Println("exit	:退出程序")
	fmt.Println("****************************************")
}

func init() {
	cmdHandlers["help"] = &cmdInfo{handler: printHelp, desc: "帮助"}
	cmdHandlers["dir"] = &cmdInfo{handler: cmdDir, desc: "当前目录子目录列表"}
	cmdHandlers["cd"] = &cmdInfo{handler: cmdCd, desc: "进入子目录或返回上层目录：cd dirname/cd .."}
	cmdHandlers["work"] = &cmdInfo{handler: cmdWorkDir, desc: "设置当前目录为工作目录"}
}

func main() {
	fmt.Println("程序路径:", exePath)
	fmt.Println("配置文件:", exePath+"cfg.ini")

	if workPath == "" {
		workPath = exePath
		curPath = workPath
	}

	printHelp(nil)

	doWork(workPath)

	OutWorkPath()
	reader := bufio.NewReader(os.Stdin)
	for {
		OutCurPath()

		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("发送错误提出程序: ", err)
			break
		}

		cmds := ParseIn(in)
		if len(cmds) == 0 {
			continue
		}

		cmd := cmds[0]
		if cmd == "exit" {
			break
		}

		info := cmdHandlers[cmd]
		if info == nil {
			if PathExists(cmd) {
				if cmd[len(cmd)-1:] != "\\" {
					cmd = cmd + "\\"
				}
				curPath = strings.Replace(cmd, "\\", "/", -1)
				workPath = curPath
				continue
			}
			fmt.Print("无效命令和参数: ", in)
			continue
		}

		info.handler(cmds[1:])
	}

	//selWorkPath()

	fmt.Println("程序结束")
	time.Sleep(time.Duration(time.Second * 3))
}
