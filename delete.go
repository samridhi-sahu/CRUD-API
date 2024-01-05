package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["roll"]
	_, err := db.Query("DELETE FROM StudentTable WHERE RollNo = ?", id)
	if err != nil {
		fmt.Println("deleting error")
		log.Fatal(err)
	}
	fmt.Println("deleted roll : ", id)
	w.WriteHeader(http.StatusNoContent)
}
