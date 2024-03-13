package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))
}

func snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet..."))
}

func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {
	//servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)

	log.Println("Starting Server on Port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
