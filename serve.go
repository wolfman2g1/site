package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("signup", signupHandler)
	http.ListenAndServe(":6060", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "test")
	f, err := os.Open("./public/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, f)

}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./public/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, f)
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse the url perameters

}
