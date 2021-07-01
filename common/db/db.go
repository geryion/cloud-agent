package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var CloudDB *gorm.DB
var DATABASE = "cloud"

func InitDB(vip string, mysqlport int) error {
	//连接数据库
	if CloudDB != nil {
		return nil
	}
	dsn := "root:root@tcp(" + vip + ":" + strconv.Itoa(mysqlport) + ")/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Failed to connect database. Err: %v", err)
	}

	//设置超时时间的上下文
	ct, _ := context.WithTimeout(context.Background(), time.Second * 5)
	//为初始化数据表的db对象添加上下文,超时退出
	tx := db.WithContext(ct)
	initUserTables(tx)
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("Failed to get sqlDB. Err: %v", err)
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(75)
	// SetConnMaxLifetime 设置连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	CloudDB = db
	return nil
}

func initUserTables(db *gorm.DB) {
	db.AutoMigrate(&User{})
}