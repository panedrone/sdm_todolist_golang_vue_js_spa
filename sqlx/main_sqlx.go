package main

import (
	"log"
	"sdm_demo_todolist/sqlx/dbal"
	"sdm_demo_todolist/sqlx/router"
)

// @title			Todolist API
// @version		0.0.1
// @description	Todolist API
// @in				header
// @BasePath		/api
// @accept			json
// @produce		json
// @host			127.0.0.1:8080
// @schemes		http
func main() {
	err := dbal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		_ = dbal.CloseDB()
	}()

	myRouter := router.New()
	// log.Fatal
	// https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
	// log.Fatal(http.ListenAndServe(":8080", myRouter))
	// https://stackoverflow.com/questions/57354389/how-to-render-static-files-within-gin-router
	log.Fatal(myRouter.Run(":8080"))
}
