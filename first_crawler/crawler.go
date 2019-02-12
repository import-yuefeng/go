package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	"github.com/djimenez/iconv-go"
	// "log"
	// "github.com/PuerkitoBio/goquery"
)

func GetJokes(){
	res, err := http.Get("http://www.4399.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	utfBody, err := iconv.NewReader(res.Body, "gb2312", "utf-8")

	// body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
 
	fmt.Println(utfBody)
}

func main(){
	GetJokes()
}
