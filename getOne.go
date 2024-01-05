package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	// storing parameters of request into map
	params := mux.Vars(r)
	id := params["roll"]
	row, err := db.Query("SELECT * FROM StudentTable WHERE RollNo = ?", id)
	if err != nil {
		fmt.Println("error in fetching this rollno info")
		log.Fatal(err)
	}
	var student Student
	if row.Next() {
		err = row.Scan(&student.RollNo, &student.FirstName, &student.LastName, &student.Class, &student.CourseName)
		if err != nil {
			fmt.Println("error in assigning this rollno info into student struct")
			log.Fatal(err)
		}
		fmt.Println("fetched row is : ", student)
	} else {
		fmt.Println("no data fetched")
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(student)
}
