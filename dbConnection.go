package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connection() {
	var err error
	password := os.Getenv("MYSQL_PASSWORD")
	db, err = sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		fmt.Println("error in connection with godb database")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("connection viability error with godb")
		log.Fatal(err)
	}
	fmt.Println("connection established with godb")
}
