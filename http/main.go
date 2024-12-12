package main

import (
	"flag"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()
}
func main() {

	http.HandleFunc("/", helloHandler)
	fmt.Println(fmt.Sprintf("Starting server at port %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
