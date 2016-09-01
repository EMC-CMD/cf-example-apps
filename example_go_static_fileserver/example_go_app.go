package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":"+p
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Printf("listening on port "+port)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("cwd: %s", wd)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my first unikernel!")
}
