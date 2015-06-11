package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w,"Hello, world!,this is a demo of GO")
  })
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}
