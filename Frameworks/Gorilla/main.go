package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
