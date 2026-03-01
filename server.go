package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", home)
	e.GET("/join", showJoin)
	e.POST("/join", join)
	e.GET("/thread/:id", showThread)
	e.POST("/thread/:id", postThread)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

type Thread struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func home(c *echo.Context) error {
	threads := []Thread{
		{ID: "dead-beef-hoge", Title: "雑談"},
		{ID: "dead-beef-foo", Title: "鉄道技術雑談"},
		{ID: "beef-dead-hoge", Title: "最新動向(スジなども)雑談"},
	}

	return c.JSON(http.StatusOK, threads)
	// スレッドの一覧とリンク、概要を出す、ログイン動線
}

func login(c *echo.Context) error {
	return c.String(http.StatusOK, "ログイン")
	// メールアドレス認証
	// 実際にメールを当該のアドレスに送ってワンタイムトークンを発行して認証する Valkeyを用いる
	// SendGrid(keyはいったん、env)の無料枠を使う
}

func showJoin(c *echo.Context) error {
	return c.String(http.StatusOK, "ユーザー登録します")
	// ユーザー登録を促す画面、入力フォーム(CSRF防止でCORS、メルアド)
}

func join(c *echo.Context) error {
	// メールアドレスに認証メールを送る(ログインと同じロジック)
	return c.String(http.StatusOK, "ユーザー登録確認メールが送信されました")
}

func showThread(c *echo.Context) error {
	threadID := c.Param("id")

	return c.String(http.StatusOK, threadID)
	// スレッドの内容を表示する
	// XSS対策は、フロントエンドでHTMLではなくテキスト(innerText)として表示することで実現
}

func postThread(c *echo.Context) error {
	threadID := c.Param("id")

	return c.String(http.StatusOK, threadID)
	// まずは文字列のみ受け付ける。将来的に画像も受け付ける
}
