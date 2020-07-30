package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type TaskList struct {
	Items []Task
}

var db *sql.DB
var err error

func createTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "content-type")

	stmt, err := db.Prepare("INSERT INTO tasks (description, status) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	rqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyval := make(map[string]string)
	json.Unmarshal(rqbody, &keyval)
	description := keyval["description"]
	status := keyval["status"]

	_, err = stmt.Exec(description, status)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New task was created")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	res, err := db.Query("SELECT id, description, status FROM tasks WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	var task Task
	for res.Next() {
		err := res.Scan(&task.Id, &task.Description, &task.Status)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE tasks SET description = ?, status = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyval := make(map[string]string)
	json.Unmarshal(body, &keyval)
	newDescription := keyval["description"]
	newStatus := keyval["status"]
	_, err = stmt.Exec(newDescription, newStatus, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task %s was updated", params["id"])
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task %s was deleted", params["id"])
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task
	var thelist TaskList

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
	thelist.Items = tasks

	json.NewEncoder(w).Encode(thelist)
}

func main() {
	db, err = sql.Open("mysql", "root:khanhdatabase@tcp(127.0.0.1:3306)/goboard")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Printf("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
