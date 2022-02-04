package main

import (
  "net/http"
  "fmt"
  "bytes"
  "io/ioutil"
  "encoding/json"
  "os"
  "unicode"
)

type Upload struct {
  Key string
}

func main() {
  var upload Upload
  destination := "https://hastebin.wiuwiu.de"
  url := destination+"/documents"

  if(len(os.Args)!=2) {
    fmt.Println("Usage: publish path/to/file")
    os.Exit(1)
  }

  content, err := ioutil.ReadFile(os.Args[1])

  if(!isASCII(string(content))){
    fmt.Println("Not an ASCII file.")
    os.Exit(2)
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &upload)
  fmt.Println(destination+"/"+upload.Key)
}

func isASCII(s string) bool {
  for _, c := range s {
    if c > unicode.MaxASCII {
      return false
    }
  }
  return true
}
