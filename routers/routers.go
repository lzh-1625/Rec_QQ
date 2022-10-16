package routers

import (
	"github.com/gin-gonic/gin"
	"Rec_QQ/controller"
)

func NewRouter() *gin.Engine{
	r:=gin.Default()
	r.POST("/",controller.Download_file)
	return r
}