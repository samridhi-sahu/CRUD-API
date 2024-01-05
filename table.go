package main

import (
	"fmt"
	"log"
)

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS StudentTable(
	RollNo     VARCHAR(50) PRIMARY KEY,
	FirstName  VARCHAR(50) NOT NULL,
    LastName   VARCHAR(50),
    Class      VARCHAR(50) NOT NULL,
    CourseName VARCHAR(50)
    )`

	_, err := db.Query(query)
	if err != nil {
		fmt.Println("error in creating StudentTable")
		log.Fatal(err)
	}
	fmt.Println("StudentTable created or existed")
}
