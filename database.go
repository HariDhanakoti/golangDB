package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "golang:password@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		var (
			id   int
			name string
		)
		rows, err := db.Query("select id, name from users where id = ?", 1)
		if err != nil {
			log.Fatal("Error ", err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Data from database " + strconv.Itoa(id) + " " + name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()
	}
}
