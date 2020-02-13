package handler

import (
	"fmt"
	"github.com/StringTek2019/go-anywhere/resolver"
	"log"
	"net/http"
	"os"
)
func ErrorWrapper(dir,prefix string,staticResolver resolver.StaticResolverType)(string,resolver.StandardResolverType){
	return "/", func(writer http.ResponseWriter, request *http.Request) {
		defer func(){
			r:=recover()
			if r!=nil{
				log.Printf("Panic:%v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		if err:=staticResolver(dir,prefix,writer,request);err!=nil{
			log.Printf("error occurred:%s",err.Error())
			if userError,ok:=err.(resolver.UserError);ok{
				switch t:=userError.(type){
				case *resolver.NoIndexPageError:
					http.Error(writer,t.Message(),http.StatusNotFound)

				case *resolver.InvalidPrefixError:
					http.Error(writer,t.Message(),http.StatusBadRequest)
				default:
					http.Error(writer,t.Message(),http.StatusInternalServerError)
				}
				return
			}
			code:=http.StatusOK
			switch{
			case os.IsNotExist(err):
				fmt.Println()
				code=http.StatusNotFound
			case os.IsPermission(err):
				code=http.StatusForbidden
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}
