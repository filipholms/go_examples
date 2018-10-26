package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Mensaje desde el metodo GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Mensaje desde el metodo post")
}

func PutUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Mensaje desde el metodo put")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Mensaje desde el metodo delete")
}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/user", GetUsers ).Methods("GET")
	r.HandleFunc("/api/user", PostUsers ).Methods("POST")
	r.HandleFunc("/api/user", PutUsers ).Methods("PUT")
	r.HandleFunc("/api/user", DeleteUsers ).Methods("DELETE")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}