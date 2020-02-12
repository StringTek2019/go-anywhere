package main

import (
	"fmt"
	"github.com/StringTek2019/go-anywhere/server"
	"os"
	"strconv"
	"strings"
)

func main() {
	start()
}

func start(){
	defer func(){
		err:=recover()
		if err!=nil{
			panic("invalid parameter")
		}
	}()
	ip:="0.0.0.0"
	port:=9999
	dir:="."
	prefix:="/"
	length:=len(os.Args)
	if length==1{
	}else if length==2{
		if os.Args[1]=="help"{
			fmt.Printf(`  go-everywhere [dir] [prefix] [ip:port]
	ip:port:default value is 0.0.0.0:9999
	dir    :the directory you want to publish.It can be absolute path and also can be relative path.if don't given,the value will be current directory
	prefix :the prefix of the url,default value is /
`)
			return
		}else{
			dir=os.Args[1]
		}
	}else if length==3{
		dir=os.Args[1]
		prefix=os.Args[2]
	}else if length==4{
		dir=os.Args[1]
		prefix=os.Args[2]
		addr:=strings.Split(os.Args[3],":")
		ip=addr[0]
		port,_=strconv.Atoi(addr[1])
	}else{
		panic("invalid parameter")
	}
	firstPos:=strings.Index(prefix,"/")
	lastPos:=strings.LastIndex(prefix,"/")
	if firstPos == -1{//have no /
		prefix= "/"+prefix+"/"
	}else if lastPos==0&&prefix!="/" {// have /
		prefix=prefix+"/"
	}else if lastPos==len(prefix)-1&&firstPos!=0{
		prefix="/"+prefix
	}
	server.Httpd(ip,port,dir,prefix)
}