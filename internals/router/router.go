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

var shouldReload = true

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

	router.GET("should-reload", func(ctx *gin.Context) {
		if shouldReload {
			ctx.JSON(http.StatusOK, gin.H{"reload": true})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"reload": false})
		}
		shouldReload = false
	})

	router.GET("/", func(ctx *gin.Context) {
		type CoffeeItem struct {
			Image  string
			Name   string
			Price  string
			Rating []bool
		}
		type PageData struct {
			Title    string
			Nickname string
			Coffees  []CoffeeItem
		}
		session := sessions.Default(ctx)

		profile := session.Get("profile")
		fmt.Println("profile", profile)
		var nickname string
		if profile == nil {
			nickname = ""
		} else {
			var ok bool
			nickname, ok = profile.(map[string]interface{})["nickname"].(string)
			if !ok {
				nickname = ""
			}
		}
		coffees := make([]CoffeeItem, 0)
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})

		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})

		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		coffees = append(coffees, CoffeeItem{
			Image:  "http://localhost:3333/public/images/coffees/Americano.png",
			Name:   "Americano",
			Price:  "R$ 9,90",
			Rating: []bool{true, true, true, true, false},
		})
		page_data := PageData{
			Title:    "Coffee Shop",
			Nickname: nickname,
			Coffees:  coffees,
		}
		fmt.Println(page_data.Nickname)
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
