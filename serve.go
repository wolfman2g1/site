package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":6060", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "test")
	f, err := io.Copy(w, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
