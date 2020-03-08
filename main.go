package main

import (
	"encoding/json"
	"io"
	"net/http"
)

//Persona escrtuctyra
type Persona struct {
	Nombre   string
	Apellido string
}

func main() {

	http.HandleFunc("/", handlerRaiz)
	http.HandleFunc("/usuarios", handlerUsuarios)

	http.ListenAndServe(":8000", nil)
}

func handlerUsuarios(response http.ResponseWriter, request *http.Request) {

	persona := Persona{"ale", "era"}

	jsResponse, err := json.Marshal(persona)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsResponse)
	// io.WriteString(response, "hola usuarios")
}

func handlerRaiz(response http.ResponseWriter, request *http.Request) {

	io.WriteString(response, "hola api")
}
