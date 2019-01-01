package main

import (
    "net/http"
	"fmt"
)

func main(){
   httpGet() 
}

func httpGet() {
    
    fmt.Print(http.Get("www.baidu.com"))

}

