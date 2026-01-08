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

func linkMXG() (*ssh.Client, error) {
	log.Printf("开始建立到墨西哥服务器连接")
	signer, errParse := ssh.ParsePrivateKey([]byte(MXGKey))
	if errParse != nil {
		log.Printf("解析私钥失败: %v", errParse)
		return nil, errParse
	}
	hostPort := fmt.Sprintf("%s:22", MXGHost)
	sshClient, errConn := ssh.Dial("tcp", hostPort, &ssh.ClientConfig{
		User:            MXGUser,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	})
	if errConn != nil {
		log.Printf("建立到墨西哥服务器连接失败: %v", errConn)
		return nil, errConn
	}
	log.Printf("完成建立到墨西哥服务器连接")
	return sshClient, nil
}

//func linkMGByProxy(client *ssh.Client) (*ssh.Client, error) {
//	log.Printf("开始建立到美国服务器连接")
//	hostPort := fmt.Sprintf("%s:22", MGHost)
//	conn, errConn := client.Dial("tcp", hostPort)
//	if errConn != nil {
//		log.Printf("通过跳板机连接到美国服务器失败: %v", errConn)
//		return nil, errConn
//	}
//	sshConn, chans, reqs, err := ssh.NewClientConn(conn, hostPort, &ssh.ClientConfig{
//		User:            MGUser,
//		Auth:            []ssh.AuthMethod{ssh.Password(MGPass)},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//		Timeout:         5 * time.Second,
//	})
//	if err != nil {
//		log.Printf("通过跳板机建立到美国服务器连接失败: %v", err)
//		return nil, err
//	}
//	sshClient := ssh.NewClient(sshConn, chans, reqs)
//	log.Printf("完成建立到美国服务器连接")
//	return sshClient, nil
//}
