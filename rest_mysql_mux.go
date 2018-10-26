package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Persona struct {
	Id_persona int `json:"id_persona"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Rut string `json:"rut"`
}


func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bodega")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	person := Persona{}

	err = db.QueryRow("SELECT * FROM persona where id_persona=2").Scan(&person.Id_persona, &person.Nombre, &person.Apellido, &person.Rut)

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	personJSON, err := json.Marshal(person)

	fmt.Print(person)

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(personJSON)



}

func main(){
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/person", GetNoteHandler).Methods("GET")

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