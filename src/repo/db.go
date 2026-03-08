package repo

import (
	"database/sql"
	"log"
	"myapp/types"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	RDB *sql.DB
}

func CreateRDB() *DataBase {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	return &DataBase{RDB: db}
}

func (db *DataBase) ThreadList() []types.Thread {
	rows, err := db.RDB.Query("SELECT thread_id, title FROM thread")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var threads []types.Thread

	for rows.Next() {
		var threadID string
		var title string
		if err := rows.Scan(&threadID, &title); err != nil {
			log.Fatal(err)
		}
		threads = append(threads, types.Thread{ID: threadID, Title: title})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return threads
}
