package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create a table
	{
		query := `
			CREATE TABLE users (
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			);
		`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	// Insert a user
	{
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id)
	}

	// Query a user
	{
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	// Query all users
	{
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user

		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}

			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	// Delete a user
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
