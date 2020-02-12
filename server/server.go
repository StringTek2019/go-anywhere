package server

import (
	"fmt"
	"github.com/StringTek2019/go-anywhere/handler"
	"github.com/StringTek2019/go-anywhere/resolver"
	"net/http"
	"os"
	"strconv"
)

/**
ip      the ip       for instance: 0.0.0.0
port    the port     for instance: 9999
dir     the local directory you want to httpd,it can be absolute path ,also can be relative path
prefix  the url prefix for instance: if the prefix='/list/' the the url must like 'http://ip:port/list/...'
 */

func Httpd(ip string,port int,dir string,prefix string){
	if validDir(dir){
		fmt.Printf("now is running on http://%s:%d%s\n",ip,port,prefix)
		http.HandleFunc(handler.ErrorWrapper(dir,prefix,resolver.StaticResolver))
		_ = http.ListenAndServe(ip+":"+strconv.Itoa(port), nil)
	}else{
		panic("the directory that you want to httpd is incorrect!")
	}

}

func validDir(dir string)bool{
	stat, err := os.Stat(dir)
	if err!=nil{
		return false
	}
	return stat.IsDir()
}