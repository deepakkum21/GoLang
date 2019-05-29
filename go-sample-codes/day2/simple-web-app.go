package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Harrison Kinsley")
}

func main() {
	fmt.Println("Server Started")
	http.HandleFunc("/", indexHandler) // path, then what function to run
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8001", nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now
}
