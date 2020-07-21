package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hudaprs/practice/gorm/models"
)

func SetHeaderContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(SetHeaderContentType)

	router.HandleFunc("/api/users", models.GetUsers).Methods("GET")
	router.HandleFunc("/api/users", models.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", models.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{id}", models.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", models.DeleteUser).Methods("DELETE")

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	models.Migration()

	handleRequests()
}
