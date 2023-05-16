package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main () {

	body := strings.NewReader("hello world")

req, _ := http.NewRequest("POST", "https://google.com", body)
client := http.Client{}
response, _ := client.Do(req)
defer response.Body.Close()

content, _ := ioutil.ReadAll(response.Body)
fmt.Println(string(content))
files, err := ioutil.ReadDir("c:\\")
if err != nil {
  log.Fatal(err)
}

for _, file := range files {
  fmt.Println(file.Name(),"\t IsDir", file.IsDir())
}


err = filepath.Walk(".",
  func(path string, info os.FileInfo, err error) error {
     if err != nil {
        return nil
     }
     fmt.Println(path, info.Size())
     return nil
  })
if err != nil {
  log.Println(err)
}

content, err = ioutil.ReadFile("c:\\9384.img")
if err != nil {
  panic(err)
}
fmt.Print(string(content))

content = []byte("Hello\nWorld\n!!!\n")
err = ioutil.WriteFile("./write-file-sample.txt", content, 0777)
if err != nil {
  panic(err)
}

}