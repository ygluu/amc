package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	ret, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return ret
}

func CutPath(pathname string) (ret string) {
	for i := len(pathname) - 1; i >= 0; i-- {
		if (pathname[i] == '\\') || (pathname[i] == '/') {
			ret = pathname[0:i]
			break
		}
	}
	ret = strings.Replace(ret, "\\", "/", -1)
	return ret + "/"
}

func CutName(pathname string) (ret string) {
	ret = ""
	for i := len(pathname) - 1; i >= 0; i-- {
		if (pathname[i] == '\\') || (pathname[i] == '/') {
			break
		}
		ret = string(pathname[i]) + ret
	}
	return
}

func ToParentDir(path string) (ret string) {
	for i := len(path) - 2; i >= 0; i-- {
		if path[i] == '/' {
			ret = path[0 : i+1]
			return
		}
	}
	return
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return (err == nil)
}

func StrToFile(s string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("File is exists: ", name)
		return
	}
	defer file.Close()
	file.WriteString(s)
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func DirList(root string) (ret []string) {
	dir, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}

	for _, fi := range dir {
		if fi.IsDir() {
			ret = append(ret, fi.Name())
		}
	}
	return
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}
