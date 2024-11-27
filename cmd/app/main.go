package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/matidev200/guardaloya-backend/internal/credentials"
	"github.com/matidev200/guardaloya-backend/internal/database"
	"github.com/matidev200/guardaloya-backend/internal/login"
	"github.com/matidev200/guardaloya-backend/internal/middleware"
	"github.com/matidev200/guardaloya-backend/internal/users"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func enableCORS(router *mux.Router) {
	FRONTEND_HOST := os.Getenv("FRONTEND_HOST")
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", FRONTEND_HOST)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	FRONTEND_HOST := os.Getenv("FRONTEND_HOST")

	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
      w.Header().Set("Access-Control-Allow-Origin", FRONTEND_HOST)
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
	router.HandleFunc("/credential", middleware.AuthMiddleware(credentials.CreateCredential)).Methods("POST")
	router.HandleFunc("/credential", middleware.AuthMiddleware(credentials.CreateCredential)).Methods(http.MethodOptions)
	router.HandleFunc("/credential/{id}", middleware.AuthMiddleware(credentials.UpdateCredential)).Methods("PATCH")
	router.HandleFunc("/credential/{id}", middleware.AuthMiddleware(credentials.GetCredential)).Methods("GET")
	router.HandleFunc("/credential/{id}", middleware.AuthMiddleware(credentials.DeleteCredential)).Methods("DELETE")
	router.HandleFunc("/credentials", middleware.AuthMiddleware(credentials.GetCredentials)).Methods("GET")

	err := http.ListenAndServe(":8080", router);
	if err != nil {
		fmt.Println("Error while starting server...", err)
	}
	fmt.Print("Server started 🤪")
}