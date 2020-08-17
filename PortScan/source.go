package main

import (
	"fmt"
	"time"
)

func say(){
	rap:=`程序编译语言：Golang 作者：Wineme

端口扫描器用于个人主机自检,网站安全防护漏扫以及内网漏扫
本扫描器代码公开,禁止用于非法用途,私自乱扫
后续添加更多功能...........

作者BILIBILI:FILTER-SNIFF`
	for _, char := range []rune(rap) {
		fmt.Printf("%s",string(char))
		time.Sleep(time.Microsecond*100000)
	}
}
func help(){
	help:=`-h 主机 (添加主机IP进行扫描)
-p 端口 (22,80 or 1-1000) 进行扫描
-v 版本信息
-say 作者留言`
	fmt.Printf(help)
}

