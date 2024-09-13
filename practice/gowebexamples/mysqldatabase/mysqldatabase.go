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
		log.Fatal(fmt.Errorf("[main] %w", err))
	}

	if err := db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("[main] %w", err))
	}

	createUsersTable(db)
	queryUsers(db)

	insertUser(db, "username1", "password1")
	queryUserById(db, 1)
	
	insertUser(db, "username2", "password2")
	queryUserById(db, 2)

	deleteUserById(db, 1)
	queryUsers(db)
}

func createUsersTable(db *sql.DB) {
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
		fmt.Println(fmt.Errorf("[createUsersTable] %w", err))
		return
	}

	fmt.Println("[createUsersTable] success")
}

func insertUser(db *sql.DB, username, password string) {
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		fmt.Println(fmt.Errorf("[insertUser] %w", err))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(fmt.Errorf("[insertUser] %w", err))
		return
	}

	fmt.Printf("[insertUser] %d inserted\n", id)
}

func queryUserById(db *sql.DB, userId int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	if err := db.QueryRow(query, userId).Scan(&id, &username, &password, &createdAt); err != nil {
		fmt.Println(fmt.Errorf("[queryUserById] %w", err))
		return
	}

	fmt.Printf("[queryUserById] %d, %s, %s, %s\n", id, username, password, createdAt.String())
}

func queryUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Println(fmt.Errorf("[queryUsers] %w", err))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Println(fmt.Errorf("[queryUsers] %w", err))
			return
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(fmt.Errorf("[queryUsers] %w", err))
		return
	}

	if len(users) < 1 {
		fmt.Println(fmt.Errorf("[queryUsers] not found"))
		return
	}

	for _, u := range users {
		fmt.Printf("[queryUsers] %d, %s, %s, %s\n", u.id, u.username, u.password, u.createdAt.String())
	}
}

func deleteUserById(db *sql.DB, userId int) {
	result, err := db.Exec(`DELETE FROM users WHERE id = ?`, userId)
	if err != nil {
		fmt.Println(fmt.Errorf("[deleteUserById] %w", err))
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(fmt.Errorf("[deleteUserById] %w", err))
		return
	}

	if rows < 1 {
		fmt.Printf("[deleteUserById] %d not found\n", userId)
		return
	}

	fmt.Printf("[deleteUserById] %d deleted\n", userId)
}
