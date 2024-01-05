package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM StudentTable")
	if err != nil {
		fmt.Println("error in fetching rows from StudentTable")
		log.Fatal(err)
	}
	fmt.Println("fetched data : ")
	var students []Student
	// iterating over fetched row from StudentTable
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.RollNo, &student.FirstName, &student.LastName, &student.Class, &student.CourseName)
		if err != nil {
			fmt.Println("error in assigning values into student struct")
			log.Fatal(err)
		}
		// inserting into slice
		students = append(students, student)
		fmt.Println(student)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(students)
}
