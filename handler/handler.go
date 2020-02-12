package handler

import (
	"github.com/StringTek2019/go-everywhere/resolver"
	"net/http"
)

func ErrorWrapper(dir,prefix string,resolver resolver.StaticResolverType)(string,string,resolver.StandardResolverType){
	return dir,prefix, func(writer http.ResponseWriter, request http.Request) {

	}
}
