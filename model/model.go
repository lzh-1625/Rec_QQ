package model

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	GroupID       int
	FileID        string `gorm:"unique"`
	FileName      string
	Busid         int
	FileSize      int
	UploadTime    int
	DeadTime      int
	ModifyTime    int
	DownloadTimes int
	Uploader      int
	UploaderName  string
}

type File struct {
	gorm.Model
	GroupID       int
	FileID        string `gorm:"unique"`
	FileName      string
	Busid         int
	FileSize      int
	UploadTime    int
	DeadTime      int
	ModifyTime    int
	DownloadTimes int
	Uploader      int
	UploaderName  string
}

type Even struct {
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	UserId      string `json:"user_id"`
	PostType    string `json:"post_type"`
	NoticeType  string `json:"notice_type"`
	File        file   `json:"file"`
	GroupId     int64  `json:"group_id"`
	Sender      sender `json:"sender"`
}

type file struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
	Url   string `json:"url"`
}

type sender struct {
	NickName string `json:"nickname"`
}
