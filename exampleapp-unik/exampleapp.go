package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
 http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
    <html><head>
    <title>CPT is awesome!</title>
    <body>
    <center>
    <br><br>
    <img src="http://www.shannonchamber.ie/wp-content/uploads/2015/09/emc-no-tag_blue_rgb.jpg">
    <br><br>
    <h1>One handsome demo app.</h1>
    <br><br>
    <h3>you will %s this demo</h3>
    </center>
    </body>
    </html>
    `, os.Getenv("foo"))
}
