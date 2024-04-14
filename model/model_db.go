package model

import (
	"fmt"
	"go-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// database入口文件

// 全局变量
var db *gorm.DB
var err error

// 连接数据库

func InitDb() {
	fmt.Printf("\n使用的数据库是%s\n\n", utils.Db)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	// 即dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("无法连接数据库：%s", err)
		return
	}

	// 用于自动迁移您的 schema，保持您的 schema 是最新的。
	err := db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()

	// 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//err = sqlDB.Close()
	//if err != nil {
	//	fmt.Println("数据库连接关闭失败！")
	//	return
	//} else {
	//	fmt.Println("数据库连接关闭")
	//}
}
