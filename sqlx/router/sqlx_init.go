package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "sdm_demo_todolist/sqlx/docs"
)

func New() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	myRouter := gin.New()

	// add swagger
	//0
	// https://github.com/swaggo/swag/issues/1019
	// You need to import docs package in your main package.
	// docs package is created by swagger when you swag init.
	//
	// https://www.youtube.com/watch?v=AtaXj2hj074
	// -->
	// https://github.com/lemoncode21/golang-crud-gin-gorm

	myRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//  https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/

	static := myRouter.Static("/static", "./static")
	// static.StaticFile("/favicon.ico", "./static/go.png")

	// === panedrone: type "http://localhost:8080" to render index.html
	static.StaticFile("/", "./static/index.html")

	//static.GET("/", func(ctx *gin.Context) {
	//	ctx.Redirect(http.StatusTemporaryRedirect, "/static")
	//})

	groupApi := myRouter.Group("/api")

	groupApi.GET("/whoiam", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Go, sqlx, sqlite3")
	})

	assignHandlers(groupApi)

	return myRouter
}

func assignHandlers(groupApi *gin.RouterGroup) {
	{
		groupProjects := groupApi.Group("/projects")
		groupProjects.GET("/", ProjectsReadAll)
		groupProjects.POST("/", ProjectCreate)
		{
			groupProject := groupProjects.Group("/:p_id")
			groupProject.GET("/", ProjectRead)
			groupProject.PUT("/", ProjectUpdate)
			groupProject.DELETE("/", ProjectDelete)
			{
				groupProjectTasks := groupProject.Group("/tasks")
				groupProjectTasks.GET("/", TasksReadByProject)
				groupProjectTasks.POST("/", TaskCreate)
			}
		}
	}
	{
		groupTask := groupApi.Group("/tasks/:t_id")
		groupTask.GET("", TaskRead)
		groupTask.PUT("", TaskUpdate)
		groupTask.DELETE("", TaskDelete)
	}
}
