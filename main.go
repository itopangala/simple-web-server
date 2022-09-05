package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParsForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request berhasil \n")
	nama := r.FormValue("nama")
	alamat := r.FormValue("alamat")
	fmt.Fprintf(w, "Nama : %s \n", nama)
	fmt.Fprintf(w, "Alamat : %s \n", alamat)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Tidak ditemukan", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method tidak didukung", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
