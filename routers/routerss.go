package routers

import (
	"github.com/gin-gonic/gin"
	"Rec_QQ/controller"
)

func NewRouters() *gin.Engine{
	r:=gin.Default()
	r.GET("/getlist",controller.FileGet)
	return r
}