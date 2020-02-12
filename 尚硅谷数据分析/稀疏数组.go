package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valnote struct {
	row int
	col int
	val int
}

func main(){

	var chessMap [11][11]int
	chessMap[1][2]=1
	chessMap[2][3]=2
	chessMap[5][3]=7
	for _,v:=range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	var sparseArr []Valnote

	valnote:=Valnote{
		row:11,
		col:11,
		val:0,
	}
	sparseArr=append(sparseArr,valnote)
	for i,v:=range chessMap{
		for j,v2 :=range v{
			if v2!=0{
				valnote:=Valnote{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr=append(sparseArr,valnote)
			}
		}
	}
	file,_:=os.OpenFile("data.txt",os.O_CREATE|os.O_RDWR,0666)
	defer file.Close()
	fmt.Println("遍历未保存数据\n")
	for i,valnote:=range sparseArr{
		var data string
		fmt.Printf("%d %d %d %d\n",i,valnote.row,valnote.col,valnote.val)
		data=fmt.Sprintf("%d %d %d\n",valnote.row,valnote.col,valnote.val)
		f:=bufio.NewWriter(file)
		f.WriteString(data)
		f.Flush()
	}
	file,_=os.Open("data.txt")
	defer file.Close()
	f:=bufio.NewReader(file)
	readdata:=make([]byte,1024*5)
	n,_:=f.Read(readdata)
	var chessMap2 [11][11]int
	readdata2:=string(readdata[:n])
	fg1:=strings.Split(readdata2,"\n")
	for i,v:=range fg1{
		if v!=""{
			if i!=0{
				fg2:=strings.Split(v," ")
				row,_:=strconv.Atoi(fg2[0])
				col,_:=strconv.Atoi(fg2[1])
				val,_:=strconv.Atoi(fg2[2])
				chessMap2[row][col]=val
			}

		}
	}
	fmt.Println("\n遍历保存读取的数据")
	for _,v:=range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}



} 