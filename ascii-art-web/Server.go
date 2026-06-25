package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("server running successfully!")
	http.ListenAndServe(":8080", nil)
}