package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matidev200/guardaloya-backend/internal/credentials"
	"github.com/matidev200/guardaloya-backend/internal/database"
	"github.com/matidev200/guardaloya-backend/internal/login"
	"github.com/matidev200/guardaloya-backend/internal/users"
)
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
      w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}
func main() {
	router := mux.NewRouter()
	enableCORS(router)
	database.NewDatabase()
	router.HandleFunc("/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/register", users.CreateUser ).Methods("POST")
	router.HandleFunc("/credential", credentials.CreateCredential).Methods("POST")
	router.HandleFunc("/credential", credentials.CreateCredential).Methods(http.MethodOptions)
	router.HandleFunc("/credential/{id}", credentials.UpdateCredential).Methods("PATCH")
	router.HandleFunc("/credential/{id}", credentials.GetCredential).Methods("GET")
	router.HandleFunc("/credential/{id}", credentials.DeleteCredential).Methods("DELETE")
	router.HandleFunc("/credentials", credentials.GetCredentials).Methods("GET")

	err := http.ListenAndServe(":8080", router);
	if err != nil {
		fmt.Println("Error while starting server...", err)
	}
	fmt.Print("Server started 🤪")
}