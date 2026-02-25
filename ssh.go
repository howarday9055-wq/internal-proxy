package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func linkBase(desc, host, user, key string) (*ssh.Client, error) {
	log.Printf("开始建立到[%s]服务器连接", desc)

	signer, errParse := ssh.ParsePrivateKey([]byte(key))
	if errParse != nil {
		log.Printf("解析私钥失败: %v", errParse)
		return nil, errParse
	}
	log.Printf("解析私钥成功")

	hostPort := fmt.Sprintf("%s:22", host)
	sshClient, errConn := ssh.Dial("tcp", hostPort, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	})
	if errConn != nil {
		log.Printf("建立到[%s]服务器连接失败: %v", desc, errConn)
		return nil, errConn
	}
	log.Printf("建立到[%s]服务器连接成功", desc)
	return sshClient, nil
}

func linkHongKong() (*ssh.Client, error) {
	return linkBase("香港", XGHost, XGUser, XGKey)
}

func linkMexico() (*ssh.Client, error) {
	return linkBase("墨西哥", MXGHost, MXGUser, MXGKey)
}

func linkSS() (*ssh.Client, error) {
	return linkBase("数数平台", SSHost, SSUser, SSKey)
}
