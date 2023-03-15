package main

import (
	"amc"
)

func main() {
	svc := &iogo.Svc{}
	svc.Exec("异步微服务示例：业务逻辑服务器", "127.0.0.1", 30000, 65535)
}
