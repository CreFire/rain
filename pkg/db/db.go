package db

import (
	"fmt"
	"github.com/CreFire/rain/config"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func Init() {
	logs := log.Default()
	log.ResetDefault(logs)
	conf := config.GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有sql
	})
	if err != nil {
		logs.Fatal("mysql not connect", log.Any("err", any(err)))
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
}

func GetDb() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Errorf("connect db server failed.")
		Init()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		Init()
	}
	return db
}

func UserSetupModel() {
	db := GetDb()
	// 自动迁移模式
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	// 创建
	db.Create(&model.User{})
}