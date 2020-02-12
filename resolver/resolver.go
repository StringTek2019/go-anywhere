package resolver

import (
	"net/http"
	"os"
)

type StaticResolverType func(dir,prefix string,writer http.ResponseWriter,request *http.Request)error
type StandardResolverType func(writer http.ResponseWriter,request http.Request)

func StaticResolver(dir,prefix string,writer http.ResponseWriter,request http.Request)error{
	path:=dir+string(os.PathSeparator)+request.URL.Path[len(prefix):]
	if file,err:=os.Open(path);err!=nil{
		return err
	}else{
	}
}