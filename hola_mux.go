hola_mux.go
package main

import (
	"fmt"
	"net/http"
)


func holaMundo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hola mundo</h1>")
}

func persona(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Juanito</h1>")
}

func perro(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Cachupin</h1>")
}

type mensaje struct {
	 msg string
}

func (m mensaje) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}



func main() {
	msg := mensaje{
		msg: "Hola Mundo",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", holaMundo)

	mux.HandleFunc("/persona", persona)

	mux.HandleFunc("/perro", perro)

	mux.Handle("/hola", msg)

	http.ListenAndServe(":8080", mux)

}