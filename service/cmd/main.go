package main

import (
	"fmt"

	"github.com/astawan/go-api-sore/core"
	"github.com/astawan/go-api-sore/service/router"
)

func main() {
	app := core.NewApp()

	env := app.Env
	gin := app.Web

	router := router.RouterConstructor(gin)
	router.NewRouter()

	gin.Run(fmt.Sprintf(":%s", env.Port))
}
