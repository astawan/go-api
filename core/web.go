package core

import "github.com/gin-gonic/gin"

func NewWeb() *gin.Engine {
	web := gin.Default()

	return web
}
