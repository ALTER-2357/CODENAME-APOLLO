package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func terrors(errors any, data string, funcname string) {

	print("error in the function " + funcname + "\n")
	print(data + "\n")
	println(errors)
	os.Exit(10)
}

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "ElSM4ZUQgGMk^*4J94#c396%L&fb9kK*s%pai0UdVqLK4C16&^Zxpta%rB"
	dbName := "unit18"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		terrors(err, "sql.Open", "dbConn")
	}
	return db
}

func main() {

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {

		router := mux.NewRouter()

		router.HandleFunc("/records/GET/", records).Methods("GET")     // MED RECORDS
		router.HandleFunc("/records/UPDATE/", records).Methods("POST") // MED RECORDS
		router.HandleFunc("/adduser/", adduser).Methods("POST")
		router.HandleFunc("/DELETE/", deleterow).Methods("POST")

		log.Fatal(http.ListenAndServe(":3031", router))

		log.Println("backends up")
		wg.Done()
	}()

	go func() {
		fs := http.FileServer(http.Dir("./site"))
		http.Handle("/", fs)

		log.Println("forntends up, go to localhost:3030 ")
		err := http.ListenAndServe(":3030", nil)
		if err != nil {
			terrors(err, "error in the website", "main ")
		}
	}()

	log.Println("backends up")

	wg.Wait()

}

func records(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nhsnumber := r.FormValue("nhsnumber")
		name := r.FormValue("name")
		gender := r.FormValue("gender")
		address := r.FormValue("address")
		phonenumber := r.FormValue("phonenumber")
		records := r.FormValue("records")
		insForm, err := db.Prepare("UPDATE apollo SET name=?, gender=?, address = ?, phonenumber =?, records=? WHERE  nhsnumber=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(name, gender, address, phonenumber, records, nhsnumber)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {

		q := "SELECT * FROM apollo WHERE nhsnumber = ?"
		nhsnumber := r.FormValue("nhsnumber")
		print(" GET, nhsnumber:" + nhsnumber)
		rows, err := db.Query(q, nhsnumber)
		if err != nil {
			terrors(err, "rows in records", "records")
		}
		jsonFile := jsonify.Jsonify(rows)
		fmt.Println(jsonFile)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(jsonFile)
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :115", "records")
		}

	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			terrors(err, "defer func", "records")
		}
	}(db)

}

func deleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nhsnumber := r.FormValue("ID")
	insForm, err := db.Prepare("DELETE FROM apollo WHERE nhsnumber=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(nhsnumber)
	if err != nil {
		terrors(err, "insForm.Exec", "deleterow")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			terrors(err, "defer func", "deleterow")
		}
	}(db)
	err = json.NewEncoder(w).Encode("done")
	if err != nil {
		terrors(err, "json.NewEncoder", "deleterow")
	}
}

func adduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	name := r.FormValue("name")
	gender := r.FormValue("gender")
	address := r.FormValue("address")
	phonenumber := r.FormValue("phonenumber")
	records := r.FormValue("records")
	insForm, err := db.Prepare("INSERT INTO apollo ( name, gender, address, phonenumber, records) VALUES(?,?,?,?,?)")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(name, gender, address, phonenumber, records)
	if err != nil {
		terrors(err, "_, err = insForm.Exec", "adduser")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			terrors(err, "defer func ", "adduser")
		}
	}(db)

	err = json.NewEncoder(w).Encode("done")
	if err != nil {
		terrors(err, "defer func in deletePost", "adduser")
	}
}
