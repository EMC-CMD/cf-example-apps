package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() { httpd() }

func httpd() {

	fmt.Println("Starting to listen!!!")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/"))))
}