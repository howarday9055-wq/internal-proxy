package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] == "" {
		log.Printf("请输入目标网络环境(xg/mxg) 和端口号")
	}

	port, errParse := strconv.Atoi(os.Args[2])
	if errParse != nil {
		log.Printf("端口号格式错误")
		return
	}
	localPort := fmt.Sprintf("127.0.0.1:%d", port)
	var (
		sshClient *ssh.Client
		err       error
	)
	switch os.Args[1] {
	case "xg":
		sshClient, err = linkXG()
	case "mxg":
		sshClient, err = linkMXG()
	default:
		log.Printf("请输入正确目标网络环境 （xg/mxg）")
		return
	}
	proxy := &httpProxy{sshClient: sshClient}
	server := &http.Server{Addr: localPort, Handler: proxy}
	log.Printf("HTTP 代理已启动，监听端口 %s", localPort)
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
