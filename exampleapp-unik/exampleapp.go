package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var count = 0

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request){
		in := r.URL.Query().Get("foo")
		w.Write([]byte(strings.ToUpper(in)))
	})
		http.HandleFunc("/reset", reset)
 http.ListenAndServe(":5000", nil)
}

func reset(w http.ResponseWriter, r *http.Request) {
	count = 0
}

func handler(w http.ResponseWriter, r *http.Request) {
	image := "https://pmcdeadline2.files.wordpress.com/2014/11/steve-jobs-pic.jpg"
	count += 1
	if count > 2 {
		image = "http://cdn.bgr.com/2015/11/bill-gates.jpg"
	}
	fmt.Fprintf(w, `
    <html><head>
    <title>Unikernels are awesome!</title>
		<meta http-equiv="refresh" content="15">
		</head>
    <body>
    <center>
    <br><br>
    <img src="%s">
    <br><br>
    <h1>Written in <a href="https://golang.org/">Go</a>.</h1>
    <br><br>
    <h3>you will %s this demo</h3>
    </center>
    </body>
    </html>
    `, image, os.Getenv("foo"))
}
