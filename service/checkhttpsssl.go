package service

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"sync"
)

//检测ssl证书，接收[]string类型，再循环单独调用检测单个域名的函数
func CheckSSL(url []string) {
	//创建一个WaitGroup 对象
	var wg sync.WaitGroup
	wg.Add(len(url))

	//循环url次数
	for i := 0; i < len(url); i++ {
		//起go程，实现并发
		go func(u string) {
			checksslOne(u)
			wg.Done()
		}(url[i])
	}
	//block，直到所有协程全部执行结束后返回
	wg.Wait()
}

// 检测单个域名的函数
func checksslOne(url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// peerCertificates,声明此变量用于接收CheckRedirect的返回
	var peerCertificates []*x509.Certificate
	client := &http.Client{
		Transport: tr,
		// CheckRedirect针对30x跳转的处理函数,非跳转的不会执行此函数
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// fmt.Printf("CheckRedirect req.Response.TLS.PeerCertificates: %#v\n", req.Response.TLS.PeerCertificates)
			peerCertificates = req.Response.TLS.PeerCertificates
			return nil
		}}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(url, " 请求失败")
		//panic(err)
	} else {
		defer resp.Body.Close()
		// 如果peerCertificates为空说明为非30x的域名,直接取TLS.PeerCertificates
		if len(peerCertificates) == 0 {
			peerCertificates = resp.TLS.PeerCertificates
		}
		certInfo := peerCertificates[0]
		fmt.Println("==========================================")
		fmt.Println("url:", url)
		fmt.Println("过期时间:", certInfo.NotAfter)
		//fmt.Println("组织信息:", certInfo.Subject)
	}
}
