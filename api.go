package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Lumodrink struct {
	Name       string `json:"Name"`
	Desc       string `json:"desc"`
	Alcoholamt int    `json:"alcoholamt"`
	Image      string `json:"image"`
	Price      int    `json:"price"`
}

var Lumodriks []Lumodrink

func returnAllLumoDrinks(w http.ResponseWriter, r *http.Request) {
	lumodriks := []Lumodrink{
		Lumodrink{Name: "Bushmills Black", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 12900},
		Lumodrink{Name: "Martell Cognac", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 27500},
		Lumodrink{Name: "Beefeater Gin", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 13600},
		Lumodrink{Name: "Jameson Irish", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 9350},
		Lumodrink{Name: "Tanqueray Gin", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 18600},
		Lumodrink{Name: "4th Street", Desc: "Test Desc", Alcoholamt: 36, Image: "https://example.com/image.jpg", Price: 4900},
	}

	fmt.Println("Endpoint Hit: returnAllLumoDrinks")
	json.NewEncoder(w).Encode(lumodriks)
}
func testPostLumoDrinks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test post Endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rest API attempt 1")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/lumodrinks", returnAllLumoDrinks).Methods("GET")
	myRouter.HandleFunc("/lumodrinks", testPostLumoDrinks).Methods("POST") // Add a new route for returning Lumodriks
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
