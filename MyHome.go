package main

import "fmt"
import "net/http"

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is the server!!")
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":9009", nil)
}
