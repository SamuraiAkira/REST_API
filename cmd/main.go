package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"REST_API/pkg"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", pkg.CreateCar).Methods("POST")
	router.HandleFunc("/cars", pkg.GetCarsList).Methods("GET")
	router.HandleFunc("/cars/{id:[0-9]+}", pkg.GetCarByID).Methods("GET")
	router.HandleFunc("/cars/{id:[0-9]+}", pkg.UpdateCarPut).Methods("PUT")
	router.HandleFunc("/cars/{id:[0-9]+}", pkg.UpdateCarPatch).Methods("PATCH")
	router.HandleFunc("/cars/{id:[0-9]+}", pkg.DeleteCar).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
