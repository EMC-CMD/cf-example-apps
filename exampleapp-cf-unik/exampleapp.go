package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `
    <html><head>
    <title>Steve Jobs is awesome!</title>
    <body>
    <center>
    <br><br>
    <img src="http://i.imgur.com/JyyOYEG.png">
    <br><br>
    <h1>One handsome demo app.</h1>
    <br><br>
    <h3>This app is tagged with: %s</h3>
    </center>
    </body>
    </html>
    `, os.Getenv("DIEGO_BRAIN_TAG"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
