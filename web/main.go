// cmd : API_HOST=localhost API_PORT=8082 NAME=John go run main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	util "github.com/verlandz/docker/web/util"
)

var port = ":8081"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello \"%v\"!", os.Getenv("NAME"))
}

func data(w http.ResponseWriter, r *http.Request) {
	url := util.StringConcat("http://", os.Getenv("API_HOST"), ":", os.Getenv("API_PORT"), "/data")
	fmt.Println("requesting to", url)

	data, ok := util.GetHttpResponse(url, map[string]string{})
	if ok {
		fmt.Fprintf(w, string(data))
	} else {
		fmt.Fprintf(w, "Fail to connect")
	}
}

func info() {
	fmt.Printf("%-8v : %v\n", "NAME", os.Getenv("NAME"))
	fmt.Printf("%-8v : %v\n", "API_HOST", os.Getenv("API_HOST"))
	fmt.Printf("%-8v : %v\n", "API_PORT", os.Getenv("API_PORT"))
	fmt.Println("Running in", runtime.Version())
	fmt.Println("Listen and serve", port)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)
	info()
	http.ListenAndServe(port, nil)
}