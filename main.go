package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	RollNo     string `json:"roll"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Class      string `json:"class"`
	CourseName string `json:"course"`
}

var db *sql.DB

func main() {
	// connecting with databse
	Connection()
	defer db.Close()
	// creating table if not existed before
	CreateTable()
	// initialized routing method
	router := mux.NewRouter()
	// handing api endpoints
	router.HandleFunc("/students", GetAll).Methods("GET")
	router.HandleFunc("/student/{roll}", GetOne).Methods("GET")
	router.HandleFunc("/student", Insert).Methods("POST")
	router.HandleFunc("/student/{roll}", Update).Methods("PUT")
	router.HandleFunc("/student/{roll}", Delete).Methods("DELETE")
	router.HandleFunc("/", ServeHome)
	// sever listening at 4000 port
	http.ListenAndServe(":4000", router)
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>this is our first crud json api</h1>"))
}
