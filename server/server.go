package server

import (
	"net/http"
	"strconv"
)

/**
ip      the ip       for instance: 0.0.0.0
port    the port     for instance: 9999
dir     the local directory you want to httpd,it can be absolute path ,also can be relative path
prefix  the url prefix for instance: if the prefix='/list/' the the url must like 'http://ip:port/list/...'
 */

func Httpd(ip string,port int,dir string,prefix string){
	http.HandleFunc()
	http.ListenAndServe(ip+":"+strconv.Itoa(port),nil)
}