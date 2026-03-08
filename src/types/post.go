package types

type Posts struct {
	ThreadID string `json:"thread_id"`
	Title    string `json:"title"`
	Total    int    `json:"total"`
	Posts    []Post `json:"posts"`
}
type Post struct {
	PostID     string       `json:"post_id"`
	UserName   string       `json:"user_name"`
	Text       string       `json:"text"`
	Time       string       `json:"time"`
	ExtContent []ExtContent `json:"ext_content"`
}

type ExtContent struct {
}
