package main

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"myapp/src/repo"
	"myapp/src/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var DB *repo.DataBase

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	DB = repo.CreateRDB()

	e.GET("/", home)
	e.GET("/join", showJoin)
	e.POST("/join", join)
	e.GET("/thread/:id", showThread)
	e.POST("/thread/:id", postThread)

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("golang.mipsparc.net"),
		// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
		Cache: autocert.DirCache("/var/www/.cache"),
		Email: "mipsparc@gmail.com",
	}

	sc := echo.StartConfig{
		Address:   ":443",
		TLSConfig: m.TLSConfig(),
	}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func home(c *echo.Context) error {
	threads := DB.ThreadList()

	return c.JSON(http.StatusOK, threads)
}

func showJoin(c *echo.Context) error {
	return c.String(http.StatusOK, "ユーザー登録します")
	// ユーザー登録を促す画面、入力フォーム(CSRF防止でCORS、メルアド)
}

func join(c *echo.Context) error {
	// 実際にメールを当該のアドレスにワンタイムトークンを発行してその画面のまま入力して認証する Valkeyを用いる
	// SendGrid(keyはいったん、env)の無料枠を使う
	return c.String(http.StatusOK, "ユーザー登録確認メールが送信されました")
}

func showThread(c *echo.Context) error {
	threadID := c.Param("id")
	// thread_id のバリデーション(大きな文字列など)

	posts := DB.GetPostsByThreadID(threadID)

	return c.JSON(http.StatusOK, posts)
}

func postThread(c *echo.Context) error {
	threadID := c.Param("id")

	// スレッド存在確認する

	p := types.TextPostInput{}
	if err := c.Bind(&p); err != nil {
		slog.Info(err.Error())
		return c.String(http.StatusBadRequest, "something went wrong")
	}
	if p.UserID == 0 ||
		p.Text == "" {
		return c.String(http.StatusBadRequest, "required fields missing")
	}

	// 制御文字除去などをする
	// あとで認証する

	p.PostID = rand.Int64()
	p.ThreadID = threadID
	jst, _ := time.LoadLocation("Asia/Tokyo")
	p.Time = time.Now().In(jst).Format(time.DateTime)

	DB.InsertTextPost(p)

	return c.String(http.StatusOK, "投稿成功(仮レスポンス)")
	// まずは文字列のみ受け付ける。将来的に画像も受け付ける
}
