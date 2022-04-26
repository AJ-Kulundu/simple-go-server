package main

import (
	"fmt"
	"net/http"
	"log"
)
func main(){
    fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting on port 8080 \n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/hello"{
		http.Error(w,"404 not Found", http.StatusNotFound)
	}
	if r.Method !="GET" {
		http.Error(w ," method not supported", http.StatusNotFound)
	}

   fmt.Fprintf(w,"hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
if err := r.ParseForm(); err!=nil {
	fmt.Fprintf(w,"Form error %v",err)
}

fmt.Fprintf(w,"Request successful \n")
name := r.FormValue("name")
fmt.Fprintf(w,"Your name is: %v",name)
}


