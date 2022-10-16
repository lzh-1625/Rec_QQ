package model

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connstring string){
	db,err:=gorm.Open("mysql",connstring)
	if err!=nil{
		panic("数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	if gin.Mode()=="release" {
		db.LogMode(false)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second*30)
	DB=db
	DB.Set("gorm:table_options","charset=utf8mb4").
	AutoMigrate(&File{})
}