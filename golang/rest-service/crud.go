package main

import (
	"net/http"
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
	"log"
)


// User structure with fields ID and Name
type User struct {
	ID   string
	Name string
}


// CreateUser function create user and generate ID for this user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        http.Error(w, http.StatusText(405), 405)
        return
	}
	
	id := r.FormValue("id")
	name := r.FormValue("name")
	if id == "" || name == "" {
        http.Error(w, http.StatusText(400), 400)
        return
	}

	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost/users")
    if err != nil {
        log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO users VALUES($1, $2)", id, name)

	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
	rowsAffected, err := result.RowsAffected()

	fmt.Fprintf(w, "User %s created successfully (%d row affected)\n", id, rowsAffected)

}
// GetUser return user from slice users by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

    id := r.FormValue("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
		return	
    }

	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost/users")
    if err != nil {
        log.Fatal(err)
	}

	usr := new(User)
	
	row := db.QueryRow("SELECT * FROM users WHERE ID=$1", id)
	err = row.Scan(&usr.ID, &usr.Name)
	if err == sql.ErrNoRows {
        http.NotFound(w, r)
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

    fmt.Fprintf(w, "%s, %s\n", usr.ID, usr.Name)
}

// UpdateUser find user in slice users by ID and modify it
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
        http.Error(w, http.StatusText(405), 405)
        return
	}
	
	id := r.FormValue("id")
	newname := r.FormValue("name")
	if id == "" || newname == "" {
        http.Error(w, http.StatusText(400), 400)
        return
	}

	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost/users")
    if err != nil {
        log.Fatal(err)
	}
	result, err := db.Exec("UPDATE users SET Name=$1 WHERE ID=$2", newname, id)
	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

    rowsAffected, err := result.RowsAffected()

    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
	}
	fmt.Fprintf(w, "User %s modified successfully (%d row affected)\n", id, rowsAffected)

}
// DeleteUser just delete user from slice by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
        http.Error(w, http.StatusText(405), 405)
        return
	}

	id := r.FormValue("id")
	if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
	}
	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost/users")
    if err != nil {
        log.Fatal(err)
	}
	result, err := db.Exec("DELETE FROM users WHERE ID=$1", id)
	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

	rowsAffected, err := result.RowsAffected()

	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
	}
	fmt.Fprintf(w, "User %s deleted successfully (%d row affected)\n", id, rowsAffected)

}
