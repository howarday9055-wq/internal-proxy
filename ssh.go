package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func linkXG() (*ssh.Client, error) {
	log.Printf("开始建立到香港服务器连接")
	signer, errParse := ssh.ParsePrivateKey([]byte(XGKey))
	if errParse != nil {
		log.Printf("解析私钥失败: %v", errParse)
		return nil, errParse
	}
	hostPort := fmt.Sprintf("%s:22", XGHost)
	sshClient, errConn := ssh.Dial("tcp", hostPort, &ssh.ClientConfig{
		User:            XGUser,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	})
	if errConn != nil {
		log.Printf("建立到香港服务器连接失败: %v", errConn)
		return nil, errConn
	}
	log.Printf("完成建立到香港服务器连接")
	return sshClient, nil
}

func linkMG() (*ssh.Client, error) {
	log.Printf("开始建立到美国服务器连接")
	hostPort := fmt.Sprintf("%s:22", MGHost)
	sshClient, errConn := ssh.Dial("tcp", hostPort, &ssh.ClientConfig{
		User:            MGUser,
		Auth:            []ssh.AuthMethod{ssh.Password(MGPass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	})
	if errConn != nil {
		log.Printf("建立到美国服务器连接失败: %v", errConn)
		return nil, errConn
	}
	log.Printf("完成建立到美国服务器连接")
	return sshClient, nil
}
