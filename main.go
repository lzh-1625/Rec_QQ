package main

import (
	"Rec_QQ/conf"
	"Rec_QQ/routers"
)

func main() {
	conf.Init()
	r:=routers.NewRouter()
	r1:=routers.NewRouters()
	r.Run(":5701")
	r1.Run("8080")
}