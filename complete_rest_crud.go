package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct{
	Title string `json:"title"`
	Description string `json:"description"`
	CreateAt time.Time `json:"create_at"`
}

var noteStore = make(map[string] Note)

var id int

func GetNoteHandler(w http.ResponseWriter, r *http.Request){
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


func PostNoteHandler(w http.ResponseWriter, r *http.Request){
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreateAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	k := vars["id"]
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteUpdate.CreateAt = note.CreateAt
		delete(noteStore, k)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

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

Syntaxis
====

/*i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}*/

	/*for i := 1; i <= 10; i++{
		print(i)
	}*/

	if 20 % 2 == 0 && 10 % 2 == 0 {
		fmt.Println("par")
	}else{
		fmt.Println("impar")
	}


var x[5]float64
	x[4] = 100
	fmt.Println(x)

var x[2] int
	x[0] = 3
	x[1] = 4

	for _, value := range x {
		println(value)
	}

x := [5]int{1, 2, 5}

	for _, value := range x{
		println(value)
	}

mapper
===
x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])



/*elements := make(map[string]string)
	elements["H"] = "hydrogen"*/

	/*elements := map[string]string{
		"H": "hydrogen",
		"B": "Boron",
	}

	name, ok := elements["H"]
	fmt.Println(ok)
	fmt.Println(name)*/

	element := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
	}

	if el, ok := element["H"]; ok{
		fmt.Println(el["name"], el["state"])
	}

FUNCIONES
===
func add(args ...int)int{
	total := 0
	for _, values := range args{
		total+= values
	}

	return total
}

func main() {

	println(add(1,2,3))

}


type Circle struct {
	x, y, r float64
}

func (c * Circle) area() float64{
	return math.Pi * c.r * c.r
}

	//var c Circle
	//c := new(Circle)
	c := Circle{0, 5, 4}
	//fmt.Println(c.x)
	fmt.Println(c.area())