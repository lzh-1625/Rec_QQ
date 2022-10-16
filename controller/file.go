package controller

import (
	"Rec_QQ/conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Data struct {
		Files []struct {
			GroupID       int    `json:"group_id"`
			FileID        string `json:"file_id"`
			FileName      string `json:"file_name"`
			Busid         int    `json:"busid"`
			FileSize      int    `json:"file_size"`
			UploadTime    int    `json:"upload_time"`
			DeadTime      int    `json:"dead_time"`
			ModifyTime    int    `json:"modify_time"`
			DownloadTimes int    `json:"download_times"`
			Uploader      int    `json:"uploader"`
			UploaderName  string `json:"uploader_name"`
		} `json:"files"`
		Folders interface{} `json:"folders"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

func FileGet(c *gin.Context) {
	response, err := http.Get("http://xcon.top:5700/get_group_root_files?group_id=" + conf.GroupID)
	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Print(err)
		return
	}
	reader := response.Body
	defer reader.Close()
	body, _ := ioutil.ReadAll(response.Body)
	finfo := FileInfo{}
	json.Unmarshal(body, &finfo)

	for i := range finfo.Data.Files {
		// fi:=model.File{}
		// fi.GroupID=finfo.Data.Files[i].GroupID
		// fi.FileID=finfo.Data.Files[i].FileID
		// fi.FileName=finfo.Data.Files[i].FileName
		// fi.Busid=finfo.Data.Files[i].Busid
		// fi.FileSize=finfo.Data.Files[i].FileSize
		// fi.UploadTime=finfo.Data.Files[i].UploadTime
		// fi.DeadTime=finfo.Data.Files[i].DeadTime
		// fi.ModifyTime=finfo.Data.Files[i].ModifyTime
		// fi.DownloadTimes=finfo.Data.Files[i].DownloadTimes
		// fi.Uploader=finfo.Data.Files[i].Uploader
		// fi.UploaderName=finfo.Data.Files[i].UploaderName
		// fmt.Print("导入第",i,"条数据……")
		// // if err:=model.DB.Create(&fi).Error;err!=nil{
		// // 	fmt.Print("数据库错误!")
		// // }
		fmt.Printf("finfo.Data.Files[i].Busid: %v\n", finfo.Data.Files[i].Busid)

	}
	c.String(200, "success")
}
