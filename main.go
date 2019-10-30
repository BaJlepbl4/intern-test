package main

import (
    "log"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	
}

func main() {
	
	
	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost/users")
    if err != nil {
        log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
        log.Fatal(err)
	} else {fmt.Printf("DB Connection established...\n")}
	

	r := mux.NewRouter()
	r.HandleFunc("/users", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/users", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}