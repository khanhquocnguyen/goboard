package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Task struct {
	Id          int
	Description string
	Status      string
}

var db *sql.DB
var err error

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task

	result, err := db.Query("SELECT id, description, status from tasks")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var task Task
		err := result.Scan(&task.Id, &task.Description, &task.Status)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

func main() {
	db, err = sql.Open("mysql", "root:khanhdatabase@tcp(127.0.0.1:3306)/goboard")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	//router.HandleFunc("/tasks", createTask).Methods("POST")
	//router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	//router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	//router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Printf("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
