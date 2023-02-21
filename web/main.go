// go run web/main.go

package main

import (
	"fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "<h2>Welcome to web dev using GO</h2>")
}
func main() {
  http.HandleFunc("/", handlerFunc)
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", nil)
}

//in browser --> http://localhost:3000/