package resolver

import (
	"io/ioutil"
	"net/http"
	"os"
)

type StaticResolverType func(dir,prefix string,writer http.ResponseWriter,request *http.Request)error
type StandardResolverType func(writer http.ResponseWriter,request *http.Request)

func StaticResolver(dir,prefix string,writer http.ResponseWriter,request *http.Request)error{
	path:=dir+string(os.PathSeparator)+request.URL.Path[len(prefix):]
	if stat, err := os.Stat(path);err!=nil{
		return err
	} else{
		if stat.IsDir(){
			filepath:=path + string(os.PathSeparator) + "index.htm"
			_, err := os.Stat(filepath)
			if err!=nil {
				filepath=path+string(os.PathSeparator)+"index.html"
				_,err =os.Stat(filepath)
				if err!=nil{
					return &NoIndexPageError{path}
				}else{
					return resolveFile(filepath,writer)
				}
			}else{
				return resolveFile(filepath,writer)
			}
		}else{
			return resolveFile(path,writer)
		}
	}
}
func resolveFile(path string,writer http.ResponseWriter)error{
	if file,err:=os.Open(path);err!=nil{
		return err
	}else{
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err!=nil{
			return err
		}
		_, err = writer.Write(content)
		if err!=nil{
			return err
		}
	}
	return nil
}
type NoIndexPageError struct{
	path string
}

func (err *NoIndexPageError) Error()string{
	return "directory ["+err.path+"] don't have an index page"
}
func (err *NoIndexPageError) Message()string{
	return "directory ["+err.path+"] don't have an index page"
}


type UserError interface{
	error
	Message()string
}