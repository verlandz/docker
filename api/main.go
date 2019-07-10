// cmd : go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

var port = ":8082"

type user struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func data(w http.ResponseWriter, r *http.Request) {
	u := []user{
		user{
			ID:   1,
			Name: "Tommy",
		},
		user{
			ID:   2,
			Name: "Mike",
		},
	}
	b, _ := json.Marshal(u)
	fmt.Fprintf(w, string(b))
}

func info() {
	fmt.Println("Running in", runtime.Version())
	fmt.Println("Listen and serve", port)
}

func main() {
	http.HandleFunc("/data", data)
	info()
	http.ListenAndServe(port, nil)
}
