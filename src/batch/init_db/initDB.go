package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/home/admin/myapp/data.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt := `
		CREATE TABLE thread (
			thread_id TEXT NOT NULL PRIMARY KEY,
			title TEXT NOT NULL
		);
		INSERT INTO thread (thread_id, title) VALUES 
			("dead-beef-hoge", "とりとめもない雑談"),
			("dead-beef-foo", "鉄道技術雑談"),
			("beef-dead-hoge", "最新動向やスジなど")
	   ;
		CREATE TABLE post (
		    post_id TEXT NOT NULL PRIMARY KEY,
		    thread_id TEXT NOT NULL,
		    user_id INT NOT NULL,
		    text TEXT NOT NULL,
		    time TEXT NOT NULL
		);
		CREATE TABLE user (
		    user_id INTEGER NOT NULL PRIMARY KEY,
		    user_name TEXT NOT NULL,
		    user_mail TEXT NOT NULL
		);
		INSERT INTO user (user_id, user_name, user_mail) VALUES
			(1, "mipsparc", "mipsparc@gmail.com"),
			(2, "yosame", "hoge@example.com")
		;
		INSERT INTO post (post_id, thread_id, user_id, text, time) VALUES
			("193959-309349", "dead-beef-hoge", 1, "テスト1 in とりとめもない雑談", "2026-03-07 00:00:00"),
			("432430-413091", "dead-beef-hoge", 1, "テスト2 in とりとめもない雑談", "2026-03-07 02:00:00"),
			("243832-093298", "dead-beef-hoge", 2, "テスト3 in yosameとりとめもない雑談", "2026-04-07 03:00:00")
		;
	`
	_, err = db.Exec(stmt)
	if err != nil {
		panic(err.Error())
	}
}
