package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/post-mime", relay)

	fmt.Println("Initialized Mailgun relay")
	http.ListenAndServe(":8080", nil)
}
func relay(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Received message")
	for k, v := range req.Header {
		fmt.Print("[", k, "]: ", v, "\n")
	}
}
