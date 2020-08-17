package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)


func PortScan(ip string,port int,result chan int,portpool *int){
	portok:=strconv.Itoa(port)
	_,err:=net.Dial("tcp",ip+":"+portok)
	if(err==nil){
		port_id:=""
		switch port {
		case 135:
			port_id=" [RCP 服务]"
		case 445:
			port_id=" [共享服务]"
		case 3389:
			port_id=" [远程桌面]"
		case 8080:
			port_id=" [Apache]"
		case 6379:
			port_id=" [Redis]"
		case 22:
			port_id=" [SSH]"


		}

		fmt.Println(ip+":"+portok+port_id+" 端口被打开")
		result<-1
	}else{
		result<-0
	}
	*portpool+=1
}

func Scan_ccc(ipaddr *string,port *string){
	resule:=make(chan int,10000)
	var portpool int
	var port1 int
	var port2 int
	fmt.Println("----------------Port Open-------------------")
	if strings.Contains(*ipaddr,"-")==true{
		ipaddr_ip:=strings.Split(*ipaddr,"-")
		ip_i4:=strings.Split(ipaddr_ip[0],".")
		ip_2i4:=strings.Split(ipaddr_ip[1],".")
		ipadd,_:=strconv.Atoi(ipaddr_ip[1])
		last_ip,_:=strconv.Atoi(ip_i4[3])
		ip_place:=ipadd-last_ip
		if len(ip_2i4)>1{
			i1,_:=strconv.Atoi(ip_i4[0])
			i2,_:=strconv.Atoi(ip_i4[1])
			i3,_:=strconv.Atoi(ip_i4[2])
			i4,_:=strconv.Atoi(ip_i4[3])
			ii1,_:=strconv.Atoi(ip_2i4[0])
			ii2,_:=strconv.Atoi(ip_2i4[1])
			ii3,_:=strconv.Atoi(ip_2i4[2])
			ii4,_:=strconv.Atoi(ip_2i4[3])
			if ii4!=255{
				ii4+=1
				i4-=1
			}else {
				ii4-=1
			}
			for{

				i4+=1
				if i4==255{
					i4=0
					i3+=1
					if i3==255{
						i3=0
						i2+=1
						if i2==255{
							i2=0
							i1+=1
						}
					}
				}
				IP:=fmt.Sprintf("%d.%d.%d.%d",i1,i2,i3,i4)
				if i1==ii1 && i2==ii2 && i3==ii3 &&i4==ii4{
					break
				}
				if strings.Contains(*port,"-")==true {
					portok := strings.Split(*port, "-")
					port1, _ = strconv.Atoi(portok[0])
					port2, _ = strconv.Atoi(portok[1])
					for i:=0;i<port2-port1;i++{
						go PortScan(IP,port1+i,resule,&portpool)
						time.Sleep(time.Microsecond*1)
					}
				}else{
					ports:=strings.Split(*port,",")
					port_num:=len(ports)
					for i := 0; i < port_num; i++ {
						port_in, _ := strconv.Atoi(ports[i])
						go PortScan(IP, port_in, resule, &portpool)
						time.Sleep(time.Microsecond*1)
					}
				}


			}
			for{
				select {
				case <-time.After(time.Second*3):
					if i1==ii1 && i2==ii2 && i3==ii3 &&i4==ii4{
						goto s
					}
				case <-resule:
				}
			}
		}else{
			if last_ip>ipadd{
				fmt.Println("IP 不规范，请重新尝试输入")
			}
		}




		if strings.Contains(*port,"-")==true{
			portok:=strings.Split(*port,"-")
			port1,_=strconv.Atoi(portok[0])
			port2,_=strconv.Atoi(portok[1])
			var much_ip int
			for ipa:=-1;ipa<ip_place;ipa++{
				ip_4:=strconv.Itoa(last_ip+ipa+1)
				ip_string:=ip_i4[0]+"."+ip_i4[1]+"."+ip_i4[2]+"."+ip_4
				for i:=0;i<port2-port1;i++{
					go PortScan(ip_string,port1+i,resule,&portpool)
					time.Sleep(time.Microsecond*1)
				}
				much_ip+=1
			}
			for{
				select {
				case <-time.After(time.Second*3):
					if(ip_place+1==much_ip){
						goto s
					}
				case <-resule:
				}
			}



		}else{
			ports:=strings.Split(*port,",")
			var much_ip int
			port_num:=len(ports)
			for ipa:=-1;ipa<ip_place;ipa++ {
				ip_4 := strconv.Itoa(last_ip + ipa + 1)
				ip_string := ip_i4[0] + "." + ip_i4[1] + "." + ip_i4[2] + "." + ip_4
				for i := 0; i < port_num; i++ {
					port_in, _ := strconv.Atoi(ports[i])
					go PortScan(ip_string, port_in, resule, &portpool)
					time.Sleep(time.Microsecond*1)
				}
				much_ip+=1
			}
			for{
				select {
				case <-time.After(time.Second*3):
					if(ip_place+1==much_ip){
						goto s
					}else{
						<-resule
					}
				}
			}

		}
	}else{
		if strings.Contains(*port,"-")==true{
			portok:=strings.Split(*port,"-")
			port1,_=strconv.Atoi(portok[0])
			port2,_=strconv.Atoi(portok[1])
			for i:=0;i<port2-port1;i++{
				go PortScan(*ipaddr,port1+i,resule,&portpool)
			}
			for{
				select {
				case <-time.After(time.Second*3):
					if(portpool==port2-port1){
						goto s
					}else{
						<-resule
					}
				}
			}
		}else{
			ports:=strings.Split(*port,",")
			port_num:=len(ports)

			for i:=0;i<port_num;i++{
				port_in,_:=strconv.Atoi(ports[i])
				go PortScan(*ipaddr,port_in,resule,&portpool)
			}
			for{
				select {
				case <-time.After(time.Second*3):
					if(portpool==port_num){
						goto s
					}else{
						<-resule
					}
				}
			}
		}
	}

	s:
		fmt.Println("----------------Scan End-------------------")
		os.Exit(0)


}

func menu(){
	var ipaddr string
	var port string
	welcome:=`Golang 高速扫描器 开发作者 BY Wineme
*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*
即刻体验Golang强大之处;;;;;;;;;;;;;;;
加群:1125581186 体验学习;;;;;;;;;;;;;
作者:Wineme --help`
	fmt.Println(welcome)
	for k,v :=range os.Args{

		switch v {
		case "-h":
			ipaddr=os.Args[k+1]
			if ipaddr==""{
				fmt.Println("请添加要扫描的主机")
				os.Exit(0)
			}
		case "-p":
			port=os.Args[k+1]
			if port==""{
				fmt.Println("请添加要扫描的端口")
				os.Exit(0)
			}
		case "-v":
			fmt.Println("Go Port Scan v1.0")
		case "-say":
			say()
		case "--help":
			help()

		}

	}
	if ipaddr!="" && port!=""{
		Scan_ccc(&ipaddr,&port)
	}
}

func main(){
	menu()
}
