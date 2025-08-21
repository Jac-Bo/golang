package app

import (
	"database/sql"

	controller "github.com/Jac-Bo/golang/controllers"
	db "github.com/Jac-Bo/golang/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	dbe := db.ConnectDB()
	a.DB = dbe
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)
	r.POST("/manga", controller.InsertManga)
	r.GET("/manga", controller.GetAllManga)
	r.GET("/manga/:id", controller.GetOneManga)
	r.PUT("/manga/:id", controller.UpdateManga)
	r.DELETE("/manga/:id", controller.DeleteManga)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
