package main

import 
(
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("server is running ...")
	http.HandleFunc("/",X)
	http.ListenAndServe(":8080",nil)
}

func X(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got request")
}