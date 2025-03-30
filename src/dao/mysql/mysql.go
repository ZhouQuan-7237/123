package mysql

import (
	"MetaCult_Insight/src/configs"
	"MetaCult_Insight/src/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql(initConfig *configs.AppConfig) (err error) {
	database := initConfig.DataBase

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.User,
		database.Pwd,
		database.Host,
		database.Port,
		database.Database)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	fmt.Println("mysql connect success")

	Db.AutoMigrate(&model.User{}, &model.UserInput{}, &model.Message{})

	return err
}
