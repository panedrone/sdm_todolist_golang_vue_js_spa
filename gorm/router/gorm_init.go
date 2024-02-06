package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	myRouter := gin.New()

	//  https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/

	static := myRouter.Static("/static", "./static")
	// static.StaticFile("/favicon.ico", "./static/go.png")

	// === panedrone: type "http://localhost:8080" to render index.html
	static.StaticFile("/", "./static/index.html")

	//static.GET("/", func(ctx *gin.Context) {
	//	ctx.Redirect(http.StatusTemporaryRedirect, "/static")
	//})

	groupApi := myRouter.Group("/api") //, func(ctx *gin.Context) { // middleware
	// ctx.Set("db", dbal.WithContext(ctx))
	//})

	groupApi.GET("/whoiam", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Go, gorm, sqlite3")
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
