package main

import 
(
	"fmt"
	//"net/http"
	"os"
)

func main()  {
	// fmt.Println("server is running ...")
	// http.HandleFunc("/",X)
	// http.ListenAndServe(":8080",nil)

	d := os.Getenv("TEST_VARIABLE_T")
	fmt.Println("TEST_VARIABLE_T RECIVED in ENV:", d)


	d2 := os.Getenv("CHAT_TAG")
	fmt.Println("CHAT_TAG RECIVED in ENV:", d2)


	d3 := os.Getenv("CHAT_SHA")
	fmt.Println("CHAT_SHA RECIVED in ENV:", d3)
}

// func X(w http.ResponseWriter, r *http.Request){
// 	fmt.Printf("got request")
// }