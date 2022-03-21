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
	dbPass := "LESf0fW1fZ3309j"
	dbName := "minerva"
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

		//student Details
		router.HandleFunc("/students/GET/", SDrecords).Methods("GET")
		router.HandleFunc("/students/UPDATE/", SDrecords).Methods("POST")
		router.HandleFunc("/students/adduser/", SDadduser).Methods("POST")
		router.HandleFunc("/students/DELETE/", SDdeleterow).Methods("POST")

		//Venue Details
		router.HandleFunc("/Venue/GET/", VDrecords).Methods("GET")
		router.HandleFunc("/Venue/UPDATE/", VDrecords).Methods("POST")
		router.HandleFunc("/Venue/adduser/", VDadduser).Methods("POST")
		router.HandleFunc("/Venue/DELETE/", VDdeleterow).Methods("POST")

		///Course Schedule Details
		router.HandleFunc("/CourseScheduleDetails/GET/", CSDrecords).Methods("GET")
		router.HandleFunc("/CourseScheduleDetails/UPDATE/", CSDrecords).Methods("POST")
		router.HandleFunc("/CourseScheduleDetails/adduser/", CSDadduser).Methods("POST")
		router.HandleFunc("/CourseScheduleDetails/DELETE/", CSDdeleterow).Methods("POST")

		//Course Details
		router.HandleFunc("/CourseDetails/GET/", CDrecords).Methods("GET")
		router.HandleFunc("/CourseDetails/UPDATE/", CDrecords).Methods("POST")
		router.HandleFunc("/CourseDetails/adduser/", CDadduser).Methods("POST")
		router.HandleFunc("/CourseDetails/DELETE/", CDdeleterow).Methods("POST")

		//Assessor Details
		router.HandleFunc("/AssessorDetails/GET/", Arecords).Methods("GET")
		router.HandleFunc("/AssessorDetails/UPDATE/", Arecords).Methods("POST")
		router.HandleFunc("/AssessorDetails/adduser/", Aadduser).Methods("POST")
		router.HandleFunc("/AssessorDetails/DELETE/", Adeleterow).Methods("POST")

		log.Fatal(http.ListenAndServe(":3031", router))

		log.Println("backends up")
		wg.Done()
	}()

	go func() {
		fs := http.FileServer(http.Dir("./site"))
		http.Handle("/", fs)

		log.Println("frontends up, go to localhost:3030 ")
		err := http.ListenAndServe(":3030", nil)
		if err != nil {
			terrors(err, "error in the website", "main ")
		}
	}()

	log.Println("backends up")

	wg.Wait()

}

/////////////////////////////////////////////////////////////////
//student Details

func SDrecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		SID := r.FormValue("SID")
		SNAME := r.FormValue("SNAME")
		address := r.FormValue("address")
		Postcode := r.FormValue("Postcode")
		email := r.FormValue("email")
		mobilenumber := r.FormValue("mobilenumber")
		CNAME := r.FormValue("CNAME")
		CID := r.FormValue("CID")
		CSD := r.FormValue("CSD")
		Achieved := r.FormValue("Achieved")
		insForm, err := db.Prepare("UPDATE Student_Details SET SNAME=?, address=?, Postcode=?, email=?, mobilenumber=?, CNAME=?, CID=?, CSD=?, Achieved=?  WHERE  SID=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(SID, SNAME, address, Postcode, email, mobilenumber, CNAME, CID, CSD, Achieved)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {

		q := "SELECT * FROM Student_Details WHERE SID = ?"
		SID := r.FormValue("SID")
		print(" GET, SID:" + SID)
		rows, err := db.Query(q, SID)
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

func SDdeleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	SID := r.FormValue("SID")
	insForm, err := db.Prepare("DELETE FROM Student_Details WHERE SID=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(SID)
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

func SDadduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	SNAME := r.FormValue("SNAME")
	address := r.FormValue("address")
	Postcode := r.FormValue("Postcode")
	email := r.FormValue("email")
	mobilenumber := r.FormValue("mobilenumber")
	CNAME := r.FormValue("CNAME")
	CID := r.FormValue("CID")
	CSD := r.FormValue("CSD")
	Achieved := r.FormValue("Achieved")
	insForm, err := db.Prepare("INSERT INTO Student_Details (SNAME, address, Postcode, email, mobilenumber, CNAME, CID, CSD, Achieved) VALUES(?,?,?,?,?,?,?,?,? )")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(SNAME, address, Postcode, email, mobilenumber, CNAME, CID, CSD, Achieved)
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

/////////////////////////////////////////////////////////////////
//Venue Details

func VDrecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		VNAME := r.FormValue("VNAME")
		VID := r.FormValue("VID")
		VADDRESS := r.FormValue("VADDRESS")
		Postcode := r.FormValue("Postcode")
		email := r.FormValue("email")
		number := r.FormValue("number")
		cost := r.FormValue("cost")
		GML := r.FormValue("GML")
		insForm, err := db.Prepare("UPDATE Venue_Details SET VNAME=?,  VADDRESS=?, Postcode=?, email=?, number=?, cost=?, GML=?  WHERE  VID=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(VNAME, VID, VADDRESS, Postcode, email, number, cost, GML)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {

		q := "SELECT * FROM Venue_Details WHERE VID = ?"
		VID := r.FormValue("VID")
		print(" GET, VID:" + VID)
		rows, err := db.Query(q, VID)
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

func VDdeleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	VID := r.FormValue("VID")
	insForm, err := db.Prepare("DELETE FROM Venue_Details WHERE VID=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(VID)
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

func VDadduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	VNAME := r.FormValue("VNAME")
	VID := r.FormValue("VID")
	VADDRESS := r.FormValue("VADDRESS")
	Postcode := r.FormValue("Postcode")
	email := r.FormValue("email")
	number := r.FormValue("number")
	cost := r.FormValue("cost")
	GML := r.FormValue("GML")

	insForm, err := db.Prepare("INSERT INTO Venue_Details (VNAME, VID, VADDRESS, Postcode, email, number, cost, GML ) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(VNAME, VID, VADDRESS, Postcode, email, number, cost, GML)
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
		terrors(err, "defer func ", "adduser")
	}
}

/////////////////////////////////////////////////////////////////
///Course Schedule Details

func CSDrecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		CNAME := r.FormValue("CNAME")
		CID := r.FormValue("CID")
		CSD := r.FormValue("CSD")
		duration := r.FormValue("duration")
		cost := r.FormValue("cost")
		insForm, err := db.Prepare("UPDATE Course_Schedule_Details SET CNAME=?, CSD=?, duration=?, cost=? WHERE  CID=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(CNAME, CID, CSD, duration, cost)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {

		q := "SELECT * FROM Course_Schedule_Details WHERE CID = ?"
		CID := r.FormValue("CID")
		print(" GET, CID:" + CID)
		rows, err := db.Query(q, CID)
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

func CSDdeleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	CID := r.FormValue("CID")
	insForm, err := db.Prepare("DELETE FROM Course_Schedule_Details WHERE CID=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(CID)
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

func CSDadduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	CNAME := r.FormValue("CNAME")
	CID := r.FormValue("CID")
	CSD := r.FormValue("CSD")
	duration := r.FormValue("duration")
	cost := r.FormValue("cost")
	insForm, err := db.Prepare("INSERT INTO Course_Schedule_Details ( CNAME, CID, CSD, duration, cost) VALUES(?,?,?,?,?)")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(CNAME, CID, CSD, duration, cost)
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

/////////////////////////////////////////////////////////////////
//Course Details
func CDrecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		CNAME := r.FormValue("CNAME")
		CID := r.FormValue("CID")
		CSD := r.FormValue("CSD")
		duration := r.FormValue("duration")
		cost := r.FormValue("cost")
		insForm, err := db.Prepare("UPDATE Course_Schedule_Details SET CNAME=?, CSD=?, duration=?, cost=? WHERE  CID=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(CNAME, CID, CSD, duration, cost)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {

		q := "SELECT * FROM Course_Details WHERE CID = ?"
		CID := r.FormValue("CID")
		print(" GET, CID:" + CID)
		rows, err := db.Query(q, CID)
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

func CDdeleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	CID := r.FormValue("CID")
	insForm, err := db.Prepare("DELETE FROM Course_Details WHERE CID=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(CID)
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

func CDadduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	CNAME := r.FormValue("CNAME")
	CID := r.FormValue("CID")
	CSD := r.FormValue("CSD")
	duration := r.FormValue("duration")
	cost := r.FormValue("cost")
	insForm, err := db.Prepare("INSERT INTO Course_Details (CNAME, CID, CSD, duration, cost) VALUES(?,?,?,?,?)")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(CNAME, CID, CSD, duration, cost)
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

/////////////////////////////////////////////////////////////////
//Assessor Details

func Arecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		ANAME := r.FormValue("name")
		AID := r.FormValue("AID")
		address := r.FormValue("address")
		Postcode := r.FormValue("Postcode")
		email := r.FormValue("email")
		number := r.FormValue("number")
		VID := r.FormValue("VID")
		insForm, err := db.Prepare("UPDATE Assessor_Details SET ANAME=?, address = ?, Postcode =?, email=?, number=?, VID=? WHERE  AID=?")
		if err != nil {
			terrors(err, "insForm", "records")
		}
		_, err = insForm.Exec(ANAME, AID, address, number, Postcode, email, VID)
		if err != nil {
			terrors(err, "insForm.Exec :92", "records")
		}

		err = json.NewEncoder(w).Encode("done")
		if err != nil {
			terrors(err, "json.NewEncoder(w).Encode(done) :97", "records")
		}
	}
	if r.Method == "GET" {
		AID := r.FormValue("VID")
		q := "SELECT * FROM Assessor_Details WHERE AID=? "
		print(" GET, AID:" + AID)
		rows, err := db.Query(q, AID)
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

func Adeleterow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	AID := r.FormValue("AID")
	insForm, err := db.Prepare("DELETE FROM Assessor_Details WHERE AID=?")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "deleterow")
	}
	_, err = insForm.Exec(AID)
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

func Aadduser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	ANAME := r.FormValue("name")
	AID := r.FormValue("AID")
	address := r.FormValue("address")
	Postcode := r.FormValue("Postcode")
	email := r.FormValue("email")
	number := r.FormValue("number")
	VID := r.FormValue("VID")
	insForm, err := db.Prepare("INSERT INTO Assessor_Details (ANAME, AID, address, number, Postcode, email, VID)  VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		terrors(err, "insForm, err := db.Prepare", "adduser")
	}
	_, err = insForm.Exec(ANAME, AID, address, number, Postcode, email, VID)
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
