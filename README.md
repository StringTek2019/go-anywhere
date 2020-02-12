# go-anywhere
a simple http server implements by golang


### Installation
go get github.com/StringTek2019/go-anywhere

### Usage
use command `go-anywhere help` to get the manual

go-everywhere \[dir] \[prefix] \[ip:port]

	ip:port:default value is 0.0.0.0:9999
	
	dir    :the directory you want to publish.It can be absolute path and also can be relative path.if don't given,the value will be current directory
	
	prefix :the prefix of the url,default value is /