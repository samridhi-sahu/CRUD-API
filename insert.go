package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	query := `INSERT INTO StudentTable(RollNo, FirstName, LastName, Class, CourseName) 
	VALUES(?,?,?,?,?)`
	_, err := db.Query(query, student.RollNo, student.FirstName, student.LastName, student.Class, student.CourseName)
	if err != nil {
		fmt.Println("error in inserting data into StudentTable")
		log.Fatal(err)
	}
	fmt.Println("inserted data is : ", student)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(student)
}
