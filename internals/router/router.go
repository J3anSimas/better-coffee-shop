package router

import (
	"coffee-shop/internals/authenticator"
	"coffee-shop/internals/handlers"
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		type PageData struct {
			Title    string
			Nickname string
		}
		session := sessions.Default(ctx)

		profile := session.Get("profile")
		fmt.Println(profile)
		nickname, ok := profile.(map[string]interface{})["nickname"].(string)
		if !ok {
			nickname = ""
		}
		page_data := PageData{
			Title:    "Coffee Shop",
			Nickname: nickname,
		}
		_ = profile
		ctx.HTML(http.StatusOK, "index.html", page_data)
	})
	router.GET("/login", handlers.Login(auth))
	router.GET("/callback", handlers.Callback(auth))
	router.GET("/logout", handlers.Logout)
	router.GET("health", func(ctx *gin.Context) {
		current_time := time.Now()
		ctx.String(http.StatusOK, "OK: "+current_time.Format(time.DateTime))
	})

	return router
}
