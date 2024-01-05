package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["roll"]
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	_, err := db.Query("UPDATE StudentTable SET RollNo = ?, FirstName = ?, LastName = ?, Class = ?, CourseName = ? WHERE RollNo = ?", student.RollNo, student.FirstName, student.LastName, student.Class, student.CourseName, id)
	if err != nil {
		fmt.Println("updating error")
		log.Fatal(err)
	}
	fmt.Println("updated data is : ", student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
