package types

type TextPostInput struct {
	UserID   int    `json:"user_id"`
	Text     string `json:"text"`
	ThreadID int
	PostID   int64
	Time     string
}
