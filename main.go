package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Bautismo struct {
	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Rut       string `json:"rut"`
	Direccion string `json:"direccion"`
}

var listaBautismos []Bautismo

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the LibroEclesial!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllBautismos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBautismos")
	json.NewEncoder(w).Encode(listaBautismos)
}

func returnSingleBautismo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, bautismo := range listaBautismos {
		if bautismo.Id == key {
			json.NewEncoder(w).Encode(bautismo)
		}
	}
}

func createNewBautismo(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var bautismo Bautismo
	json.Unmarshal(reqBody, &bautismo)
	// update our global Articles array to include
	// our new Article
	listaBautismos = append(listaBautismos, bautismo)

	json.NewEncoder(w).Encode(bautismo)
}

func deleteBautismo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, bautismo := range listaBautismos {
		if bautismo.Id == id {
			listaBautismos = append(listaBautismos[:index], listaBautismos[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/bautismos", returnAllBautismos)
	myRouter.HandleFunc("/bautismo", createNewBautismo).Methods("POST")
	myRouter.HandleFunc("/bautismo/{id}", deleteBautismo).Methods("DELETE")
	myRouter.HandleFunc("/bautismo/{id}", returnSingleBautismo)
	log.Fatal(http.ListenAndServe(":10003", myRouter))
}

func main() {
	listaBautismos = []Bautismo{
		Bautismo{Id: "1", Nombre: "David", Rut: "131313-K", Direccion: "Avenida Central 123"},
		Bautismo{Id: "2", Nombre: "Nicolas", Rut: "130220-1", Direccion: "Pasaje 321"},
		Bautismo{Id: "3", Nombre: "Carlos", Rut: "1321321-4", Direccion: "Edicio 342"},
	}
	handleRequests()
}
