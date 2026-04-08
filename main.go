package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入目标网络环境(hk/mx/nhk, 默认 hk): ")
	envInput, _ := reader.ReadString('\n')
	env := strings.TrimSpace(envInput)
	if env == "" {
		env = "hk"
	}

	fmt.Print("请输入本地监听端口(默认 8085): ")
	portInput, _ := reader.ReadString('\n')
	portStr := strings.TrimSpace(portInput)
	if portStr == "" {
		portStr = "8085"
	}

	port, errParse := strconv.Atoi(portStr)
	if errParse != nil {
		log.Printf("端口号格式错误: %v", errParse)
		return
	}
	localPort := fmt.Sprintf("127.0.0.1:%d", port)

	var (
		sshClient *ssh.Client
		err       error
	)
	switch env {
	case "hk":
		sshClient, err = linkHongKong()
	case "mx":
		sshClient, err = linkMexico()
	case "nhk":
		sshClient, err = linkNewHongKong()
	default:
		log.Printf("请输入正确目标网络环境(hk/mx/nhk)")
		return
	}
	proxy := &httpProxy{sshClient: sshClient}
	server := &http.Server{Addr: localPort, Handler: proxy}
	log.Printf("HTTP 代理已启动，监听端口 %s", localPort)
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
