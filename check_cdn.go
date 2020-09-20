package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func iscdn(domain chan string,ret chan string){
	for{
		domain:=<-domain
		cmd:=exec.Command("cmd.exe","/c","nslookup "+domain)
		res,err:=cmd.Output()
		if(err!=nil){
			fmt.Println("errr in cmd.exe not power",err)
		}
		canary:=string(res)
		reg:=regexp.MustCompile("(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}")
		str:=reg.FindAllString(canary,-1)
		if len(str)>2{
			ret<- domain+",CDN存在"
		}else{

			ret<- domain+",CDN不存在"
		}
	}

}

func open_check(filenam string) []string{
	file,err:=os.Open(filenam)
	if err!=nil{
		fmt.Println("not this file")
		os.Exit(0)
	}
	str:=make([]byte,1024*5000)
	n,_:=file.Read(str)
	split:=string(str[:n])
	strok:=strings.Split(split,"\r\n")
	return strok
}

func say(){
	rap:=`开发作者：Wineme
任何语言都没有土拨鼠的速度快。
那么你应该体验一下Golang的魅力。
即可起.....体验一下并行的的快速吧。
10
9
8
7
6
5
4
3
2
1
`
	for _, char := range []rune(rap) {
		fmt.Printf("%s",string(char))
		time.Sleep(time.Microsecond*100000)
	}
}

func producer(urllist []string){
	urllistx:=make(chan string,1024*5000)
	channer:=make(chan string,1024*50)
	say()
	for _,v :=range urllist{
		if v!=""{
			urllistx<-v
		}
	}
	for i:=0;i<30;i++{
		go iscdn(urllistx,channer)
	}
	id:=0
	os.OpenFile("result.csv",os.O_WRONLY|os.O_CREATE,0666)
	for{
		select {
		case <-time.After(time.Second*3):
			os.Exit(0)
		case res:=<-channer:
			id=id+1
			str:=strconv.Itoa(id)+","+res
			fmt.Println(str)
			file,_:=os.OpenFile("result.csv",os.O_WRONLY|os.O_APPEND,0666)
			file.WriteString(str+"\r\n")
			file.Close()
		}

	}
}

func main(){
	if len(os.Args)==1{
		fmt.Println("i am sorry,you need enter the filename!")
	}else{
		filename:=os.Args[1]
		check_url:=open_check(filename)
		producer(check_url)
	}


}
