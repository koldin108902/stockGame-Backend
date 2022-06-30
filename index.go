package main

import (
	"net/http"
	"stockServer/router"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	app := httprouter.New()

	go func() {
		router.UpdateStock()
		time.Sleep(time.Minute * 1)
	}()

	// user
	app.GET("/user/info/:userId", router.GetUserInfo)
	app.POST("/user/my", router.GetMyInfo)

	// stock
	app.GET("/stock/all", router.AllStockList)
	app.POST("/stock/buy", router.BuyStock)
	app.POST("/stock/sell", router.SellStock)

	http.ListenAndServe(":8080", app)
}
