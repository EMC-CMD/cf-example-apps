package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-martini/martini"
)

var path = "/data/data.json"

type information []string

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		data, err := readFile()
		if err != nil {
			data = []string{}
			data = append(data, "file could not be read")
			data = append(data, "error: "+err.Error())
		}

		return mainPage(data)
	})
	m.Get("/submit", func(req *http.Request) string {
		key := req.URL.Query().Get("firstname")
		val := req.URL.Query().Get("lastname")
		data, _ := readFile()
		data = append(data, key+" "+val)
		dataString, _ := json.Marshal(data)
		err := writeFile(string(dataString))
		if err != nil {
			data = []string{}
			data = append(data, "file could not be read")
			data = append(data, "error: "+err.Error())
		}
		return redirect(key + " " + val)
	})
	m.RunOnAddr(":" + os.Getenv("PORT"))

}

func readFile() (information, error) {
	var readInfo information
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return readInfo, err
	}
	err = json.Unmarshal(data, &readInfo)
	if err != nil {
		return readInfo, err
	}
	return readInfo, nil
}

func writeFile(data string) error {
	err := ioutil.WriteFile(path, []byte(data), 0777)
	if err != nil {
		err := os.MkdirAll(filepath.Dir(path), 0777)
		if err != nil {
			return err
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func mainPage(info information) string {
	infoString := ""
	for _, entry := range info {
		infoString += entry + "<br>"
	}
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>

<form action="/submit">
  Write something:<br>
  <input type="text" name="firstname" value="dev">
  <br>
  Write something else:<br>
  <input type="text" name="lastname" value="ops">
  <br><br>
  <input type="submit" value="Submit">
</form>

<p>Saved data: <br>%s</p>

</body>
</html>
`, infoString)
}

func redirect(writtenData string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
   <!-- HTML meta refresh URL redirection -->
   <meta http-equiv="refresh"
   content="1; url=/">
</head>
<body>
   <p>Added '%s' to file!</p>
</body>
</html>`, writtenData)
}
