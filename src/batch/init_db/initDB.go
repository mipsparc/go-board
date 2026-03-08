package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = os.Remove("./data.db")

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := `
		CREATE TABLE threads (
			thread_id TEXT NOT NULL PRIMARY KEY,
			title TEXT NOT NULL
		);
	`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt = `
		INSERT INTO threads (thread_id, title) VALUES 
			("dead-beef-hoge", "とりとめもない雑談"),
			("dead-beef-foo", "鉄道技術雑談"),
			("beef-dead-hoge", "最新動向やスジなど")
	   ;
	`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
