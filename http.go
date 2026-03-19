package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net/http"
)

type httpProxy struct {
	sshClient *ssh.Client
}

func (p *httpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		p.handleHTTPS(w, r)
	} else {
		p.handleHTTP(w, r)
	}
}

func (p *httpProxy) handleHTTPS(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTPS --> %s %s", r.Method, r.Host)
	destConn, err := p.sshClient.Dial("tcp", r.Host)
	if err != nil {
		http.Error(w, "无法通过 SSH 连接目标: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer func() { _ = destConn.Close() }()

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "服务器不支持 Hijacking", http.StatusInternalServerError)
		return
	}
	clientConn, _, errHijack := hijacker.Hijack()
	if errHijack != nil {
		http.Error(w, errHijack.Error(), http.StatusServiceUnavailable)
		return
	}
	defer func() { _ = clientConn.Close() }()
	_, _ = clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	go io.Copy(destConn, clientConn)
	io.Copy(clientConn, destConn)
}

func (p *httpProxy) handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP --> %s %s", r.Method, r.URL.String())

	transport := &http.Transport{
		DialContext: p.sshClient.DialContext,
	}

	outReq := new(http.Request)
	*outReq = *r
	outReq.RequestURI = ""

	resp, err := transport.RoundTrip(outReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
