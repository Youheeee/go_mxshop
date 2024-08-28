package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"mxshop_srvs/user_srv/model"
)

//md5加密
//func genMd5(code string) string {
//	Md5 := md5.New()
//	_, _ = io.WriteString(Md5, code)
//	return hex.EncodeToString(Md5.Sum(nil))
//}

func main() {
	dsn := "root:123456@tcp(118.25.150.137:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	// 在这里使用 db 进行数据库操作
	_ = db.AutoMigrate(&model.User{})

}
