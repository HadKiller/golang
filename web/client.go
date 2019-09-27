package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func LoginIn(respon http.ResponseWriter,request *http.Request)  {
	e := request.ParseForm()
	if e!=nil{
		log.Fatal(e)
	}
	if request.Method==http.MethodGet{
		files, _ := template.ParseFiles("web/View/LoginPage.html")
		fmt.Fprintf(respon, "Hello world!")
		files.Execute(respon,nil)
	} else {
		fmt.Println("username:", request.Form["username"])
		fmt.Println("password:", request.Form["password"])

	}
}

func main() {
	http.HandleFunc("/SayHi",SayHello)
	http.HandleFunc("/Login",LoginIn)
	e := http.ListenAndServe(":9099", nil)
	if e!=nil{
		log.Fatal("Listen and Server:",e)
	}

}

func SayHello(respon http.ResponseWriter,request *http.Request)  {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path:",request.URL.Path)
	fmt.Println("scheme:",request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(respon, "Hello world!")
}
