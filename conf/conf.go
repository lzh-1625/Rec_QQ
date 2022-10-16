package conf

import (
	// "Rec_QQ/model"
	"fmt"
	"os"

	// "strings"

	"github.com/go-ini/ini"
)

var (
	HttpPort   string
	SendPort   string
	GroupID    string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	CqIP       string
	CqPort     string
)

func Init() {
	os.Mkdir("./Downloads", os.ModePerm)
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Print(err)
	}
	LoadConf(file)

	// path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	// model.Database(path)
}

func LoadConf(file *ini.File) {
	HttpPort = file.Section("service").Key("HttpPort").String()
	SendPort = file.Section("service").Key("SendPort").String()
	GroupID = file.Section("service").Key("GroupId").String()
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	CqIP = file.Section("cqhttp").Key("IP").String()
	CqPort = file.Section("cqhttp").Key("Port").String()
}
