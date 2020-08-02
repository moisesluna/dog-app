package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//Dog Model
type Dog struct {
	ID int `json:"ID"`

	Name string `json:"Name"`

	Breed string `json:"Breed"`
}

var database *sql.DB

func getDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, _ := database.Query("SELECT * FROM dog")

	var dogs = []Dog{}
	var id int
	var name string
	var breed string
	for rows.Next() {
		rows.Scan(&id, &name, &breed)
		dog := Dog{}
		dog.ID = id
		dog.Name = name
		dog.Breed = breed
		dogs = append(dogs, dog)
	}
	json.NewEncoder(w).Encode(dogs)
}

func createDog(w http.ResponseWriter, r *http.Request) {
	var dog Dog
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &dog)
	statement, _ := database.Prepare("INSERT INTO dog (name, breed) VALUES (?, ?)")
	statement.Exec(dog.Name, dog.Breed)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dog)

}

func updateDog(w http.ResponseWriter, r *http.Request) {
	var dog Dog
	vars := mux.Vars(r)
	dogID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}

	json.Unmarshal(reqBody, &dog)

	statement, _ := database.Prepare("UPDATE dog SET name = ?, breed = ? where id = ?")
	statement2, err := statement.Exec(dog.Name, dog.Breed, dogID)
	fmt.Println(statement2)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dog)

}

func getOneDog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dogID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	rows, _ := database.Query(fmt.Sprintf("SELECT * FROM dog where id = %d", dogID))
	var id int
	var name string
	var breed string
	dog := Dog{}
	for rows.Next() {
		rows.Scan(&id, &name, &breed)
		dog.ID = id
		dog.Name = name
		dog.Breed = breed
	}
	json.NewEncoder(w).Encode(dog)
}

func deleteDog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dogID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	stmt, _ := database.Prepare("DELETE FROM dog where id = ?")
	stmt2, err := stmt.Exec(dogID)
	fmt.Println(stmt2)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "The task with ID %v has been remove successfully", dogID)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	database, _ = sql.Open("sqlite3", "./moises.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS dog (id INTEGER PRIMARY KEY, name TEXT, breed TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO dog (name, breed) VALUES (?, ?)")
	statement.Exec("Desmond", "Labrador")

	router.HandleFunc("/", getDogs)
	router.HandleFunc("/dogs", getDogs).Methods("GET")
	router.HandleFunc("/dogs", createDog).Methods("POST")
	router.HandleFunc("/dogs/{id}", getOneDog).Methods("GET")
	router.HandleFunc("/dogs/{id}", deleteDog).Methods("DELETE")
	router.HandleFunc("/dogs/{id}", updateDog).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
