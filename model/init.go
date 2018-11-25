package model

import (
	"fmt"
	"holiday/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DataBase struct {
	DB *gorm.DB
}

var PostgresDB *DataBase

func open(
	dbName string,
	user string,
	password string,
	port int,
	sslMode string,
	host string,
) *gorm.DB {
	connectInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)
	connect, err := gorm.Open("postgres", connectInfo)
	var dbConnect = logger.Info{
		Package: "model",
		Action:  "db connect",
		Message: "db connect fail",
	}
	if err != nil {
		fmt.Println(err)
		logger.PanicLog(dbConnect)
		return nil
	}
	connect.LogMode(true)
	connect.DB().SetMaxIdleConns(0)
	return connect
}

func initDB() *gorm.DB {
	return open(
		"tb_holidays",
		"postgres",
		"root",
		5432,
		"disable",
		"127.0.0.1",
	)
}

func GetDB() *gorm.DB {
	return initDB()
}

func (d *DataBase) Init() {
	PostgresDB = &DataBase{
		DB: GetDB(),
	}
}

func (d *DataBase) Close() {
	PostgresDB.DB.Close()
}
