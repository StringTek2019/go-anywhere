package handler

import (
	"github.com/StringTek2019/go-anywhere/resolver"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)



func notExistError(_,_ string,_ http.ResponseWriter,_ *http.Request)error{
	return os.ErrNotExist
}
func permissionDeniedError(_,_ string,_ http.ResponseWriter,_ *http.Request)error{
	return os.ErrPermission
}
func InternalServerError(_,_ string,_ http.ResponseWriter,_ *http.Request)error{
	panic("500")
}
func noIndexPageError(_,_ string,_ http.ResponseWriter,_ *http.Request)error{
	return &resolver.NoIndexPageError{Path:"."}
}
func invalidPrefixError(_,_ string,_ http.ResponseWriter,_ *http.Request) error{
	return &resolver.InvalidPrefixError{Prefix:"/list/"}
}


var data=[]struct{
	resolver resolver.StaticResolverType
	code int
	msg string
}{
	{notExistError, 404, "Not Found"},
	{permissionDeniedError, 403, "Forbidden"},
	{InternalServerError, 500, "Internal Server Error"},
	{noIndexPageError, 404, "directory [.] don't have an index page"},
	{invalidPrefixError,400,"url must start with /list/"},

}


func TestErrorWrapper(t *testing.T) {
	for _,tt:=range data{
		resp:=httptest.NewRecorder()
		request:=httptest.NewRequest(http.MethodGet,"http://127.0.0.1",nil)
		_,solver:=ErrorWrapper(".","/",tt.resolver)
		solver(resp,request)
		verifyResponse(resp.Result(),tt.code,tt.msg,t)
	}
}

func TestHttpServer(t *testing.T){
	for _,tt:=range data{
		_,solver:=ErrorWrapper(".","/",tt.resolver)
		server:=httptest.NewServer(http.HandlerFunc(solver))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp,tt.code,tt.msg,t)
	}

}

func verifyResponse(resp *http.Response,expectedCode int,expectedMsg string,t *testing.T){
	message,_:=ioutil.ReadAll(resp.Body)
	msg:=strings.Trim(string(message),"\n")
	if resp.StatusCode!=expectedCode||string(msg)!=expectedMsg{
		t.Errorf("Excepted (Code:%d,Msg:%s) but got (Code:%d,Msg:%s)",expectedCode,expectedMsg,resp.StatusCode,msg)
	}
}