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

	d := os.Getenv("CHAT_COMMIT")
	fmt.Println("CHAT_COMMIT RECIVED in ENV:", d)


	d2 := os.Getenv("CHAT_TAG")
	fmt.Println("CHAT_TAG RECIVED in ENV:", d2)
}

// func X(w http.ResponseWriter, r *http.Request){
// 	fmt.Printf("got request")
// }