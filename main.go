package main

import 
(
	"fmt"
	"net/http"
	"os"
)

func main()  {
	// fmt.Println("server is running ...")
	// http.HandleFunc("/",X)
	// http.ListenAndServe(":8080",nil)

	d := os.Getenv("CHAT_TAG")
	fmt.Println("RECIVED in ENV:", d)
}

func X(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got request")
}