package repo

import (
	"database/sql"
	"log/slog"
	"myapp/src/types"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	RDB *sql.DB
}

func CreateRDB() *DataBase {
	db, err := sql.Open("sqlite3", "/home/admin/myapp/data.db")
	if err != nil {
		panic(err.Error())
	}

	return &DataBase{RDB: db}
}

func (db *DataBase) ThreadList() []types.Thread {
	rows, err := db.RDB.Query("SELECT thread_id, title FROM thread")
	if err != nil {
		slog.Error("Error getting thread list", err.Error())
		return []types.Thread{}
	}
	defer rows.Close()

	var threads []types.Thread

	for rows.Next() {
		var threadID string
		var title string
		if err := rows.Scan(&threadID, &title); err != nil {
			slog.Error("Error getting thread list")
			return []types.Thread{}
		}
		threads = append(threads, types.Thread{ThreadID: threadID, Title: title})
	}
	err = rows.Err()
	if err != nil {
		slog.Error("Error getting thread list")
		return []types.Thread{}
	}

	return threads
}

func (db *DataBase) GetPostsByThreadID(threadID string) types.Posts {
	rows, err := db.RDB.Query(`
		SELECT post_id, thread.title, text, time
		FROM post
		JOIN thread ON post.thread_id = thread.thread_id
		WHERE post.thread_id=?
	`, threadID)
	if err != nil {
		slog.Error("Error getting post list", err.Error())
		return types.Posts{}
	}

	defer rows.Close()
	// 0件だったときのパターン用意

	var title string

	var postList []types.Post
	for rows.Next() {
		var postID string
		var text string
		var time string
		if err := rows.Scan(&postID, &title, &text, &time); err != nil {
			slog.Error("Error getting post list")
			return types.Posts{}
		}
		postList = append(postList,
			types.Post{
				PostID:     postID,
				Text:       text,
				Time:       time,
				ExtContent: nil,
			},
		)
	}

	posts := types.Posts{
		ThreadID: threadID,
		Title:    title,
		Total:    9999,
		Posts:    postList,
	}

	return posts
}
