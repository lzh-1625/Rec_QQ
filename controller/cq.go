package controller

import (
	"Rec_QQ/conf"
	"Rec_QQ/model"
	"Rec_QQ/pkg"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Download_file(c *gin.Context) {
	message := model.Even{}
	// message :=make(map[string]interface{})
	c.BindJSON(&message)
	if message.PostType == "notice" {
		if message.NoticeType == "group_upload" {
			GID := strconv.FormatInt(message.GroupId, 10)
			fmt.Printf("GID: %v\n", GID)
			fmt.Printf("conf.GroupID: %v\n", conf.GroupID)
			if GID == conf.GroupID {
				fmt.Println("file uploading")
				gi := (message.GroupId)
				if message.GroupId == gi {
					url := message.File.Url
					name := message.File.Name
					go pkg.Download(url, name)
				}
			}
		}
	}
}
