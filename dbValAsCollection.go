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

		rows, _ := db.Query("select * from users")

		type User struct {
			name string
			id   int
		}
		users := []User{}

		for rows.Next() {
			var user User
			err := rows.Scan(&user.id, &user.name)
			if err != nil {
				log.Fatal("Scan : %v", err)
			}
			users = append(users, user)
		}

		for _, v := range users {
			log.Println("id " + strconv.Itoa(v.id) + " Name " + v.name)
		}
		log.Println("End of program")

		defer db.Close()
	}
}
