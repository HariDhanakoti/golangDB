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

		stmt, _ := db.Prepare("select id, name from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()
		rows, err := stmt.Query(13)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var (
			id   int
			name string
		)

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal("Scan : %v", err)
			}
		}
		log.Println("id " + strconv.Itoa(id) + " Name " + name)

		log.Println("End of program")

		defer db.Close()
	}
}
