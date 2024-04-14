package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

// 全局变量
var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone        int
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

// initialize the file: "config.ini"
func init() {
	// 检查环境变量是否设置，如果没有设置，则使用默认路径
	configFile := os.Getenv("GO_BLOG_CONFIG_FILE")
	if configFile == "" {
		configFile = "config/config.ini" // 默认路径
	}
	file, err := ini.Load(configFile)
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	// 读取 config.ini 中的各个section
	LoadServer(file)
	LoadDatabase(file)
	LoadQiniu(file)
}

// MustString()的意思是：如果Key对应的值为nil，那么就赋值默认

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("HttpPort").MustString("dxp8964zzp")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("pzh")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("1113")
	DbName = file.Section("database").Key("DbName").MustString("goblog")
}

func LoadQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
