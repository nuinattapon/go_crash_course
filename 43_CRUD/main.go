package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Define mysqlDB as a global variable
var mysqlDB *sqlx.DB

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sqlx.DB) {

	if mysqlDB == nil {
		dbDriver := "mysql"
		dbUser := "dev"
		dbPass := "Welcome1"
		dbName := "nui"
		dbHost := "tcp(192.168.1.6:3306)"
		dbSource := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?parseTime=true"
		// "dev:Welcome1@tcp(192.168.1.6:3306)/nui?parseTime=true"
		db, err := sqlx.Open(dbDriver, dbSource)
		if err != nil {
			panic(err.Error())
		} else {
			log.Println("Database connection is initialized")
		}
		db.SetMaxOpenConns(20)
		mysqlDB = db
	}
	return mysqlDB
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	res := []Employee{}
	err := db.Select(&res, "SELECT * FROM employee ORDER BY id DESC LIMIT 100")
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	emp := Employee{}
	err := db.Get(&emp, "SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	emp := Employee{}
	err := db.Get(&emp, "SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO employee(name, city) VALUES(?,?)", name, city)
		tx.Commit()
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		tx := db.MustBegin()
		tx.MustExec("UPDATE employee SET name=?, city=? WHERE id=?", name, city, id)
		tx.Commit()
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("id")
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM employee WHERE id=?", id)
	tx.Commit()
	log.Println("DELETE")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
