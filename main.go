package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/post-mime", relay)

	fmt.Println("Initialized Mailgun relay")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func relay(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Received message")
	for k, v := range req.Header {
		fmt.Print("[", k, "]: ", v, "\n")
	}
	fmt.Println("[BODY]:")
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
