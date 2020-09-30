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

func iscdn(domain string) string {
	for{
		cmd:=exec.Command("nslookup",domain)
		res,err:=cmd.Output()
		if(err!=nil){
			fmt.Println("errr in cmd.exe not power",err)
		}
		canary:=string(res)
		if len(canary)<70{
			return domain
		}
		reg:=regexp.MustCompile("(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}")
		str:=reg.FindAllString(canary,-1)
		if strings.Contains(canary,"Aliases")==true{
			return domain+",HAVE CDN"
		}else{
			if len(str)>2{
				return domain+",HAVE CDN"
			}else{
				if len(str)>1{
					return domain+",NOT CDN,"+str[1]
				}
			}
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
		time.Sleep(time.Microsecond*1000)
	}
}

func producer(urllist []string){
	say()
	id:=0
	os.OpenFile("result.txt",os.O_WRONLY|os.O_CREATE,0666)
	for _,v :=range urllist{
		if v!=""{
			id=id+1
			str:=strings.Replace(v,"http://","",-1)
			str2:=strings.Replace(str,"https://","",-1)
			str3:=strings.Replace(str2,"/","",-1)
			res:=iscdn(str3)
			report:=strconv.Itoa(id)+":"+res
			fmt.Println(report)
			file,_:=os.OpenFile("result.txt",os.O_WRONLY|os.O_APPEND,0666)
			file.WriteString(report+"\r\n")
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
