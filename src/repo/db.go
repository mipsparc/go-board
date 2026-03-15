package repo

import (
	"database/sql"
	"log/slog"
	"myapp/src/types"
	"strings"

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
		var threadID int
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

func (db *DataBase) GetPostsByThreadID(threadID int) types.Posts {
	rows, err := db.RDB.Query(`
		SELECT post_id, thread.title, user_name, text, time
		FROM post
		JOIN thread ON post.thread_id = thread.thread_id
		JOIN user ON post.user_id = user.user_id
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
		var userName string
		var text string
		var time string
		if err := rows.Scan(&postID, &title, &userName, &text, &time); err != nil {
			slog.Error("Error getting post list")
			return types.Posts{}
		}
		postList = append(postList,
			types.Post{
				PostID:     postID,
				UserName:   userName,
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

func (db *DataBase) CheckExistenceThread(threadID int) bool {
	result, err := db.RDB.Query(`
		SELECT COUNT(1)
		FROM thread
		WHERE post.thread_id=?
		`, threadID)
	if err != nil {
		slog.Error("Error checking existence thread", err.Error())
	}
	// nilになることは考慮外だが、静的解析避け
	if result == nil {
		return false
	}
	defer result.Close()

	if !result.Next() {
		return false
	}
	return true
}

func (db *DataBase) InsertTextPost(post types.TextPostInput) {
	_, err := db.RDB.Exec(strings.ReplaceAll(`
			INSERT INTO post (post_id, thread_id, user_id, text, time) VALUES
			   (?, ?, ?, ?, CURRENT_TIMESTAMP)
			ON DUPLICATE KEY UPDATE ""post_id"" = ""post_id"";
		`, `""`, "`"),
		post.PostID, post.ThreadID, post.UserID, post.Text)
	if err != nil {
		slog.Error("Error inserting post", err.Error())
		return
	}
}
