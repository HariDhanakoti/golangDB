package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "golang:password@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Success")

		var queryName string
		var id int
		err := db.QueryRow("select id, name from users where id = ?", 1).Scan(&id, &queryName)
		if err != nil {
			log.Fatal("Error ", err)
			if err == sql.ErrNoRows {
				log.Println("No rows returned from database")
			}
		}
		log.Println("Id is " + strconv.Itoa(id) + " Name is " + queryName)
		defer db.Close()
	}
}
