package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Env   *Env
	Mysql *gorm.DB
	Web   *gin.Engine
}

func NewApp() *Application {
	app := Application{}
	app.Env = NewEnv()
	app.Mysql = NewMySQLConnection(&app)
	app.Web = NewWeb()

	return &app
}
