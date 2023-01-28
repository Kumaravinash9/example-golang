package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	url := "root" + ":" + "Bhai@vi9" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "dhandolat" + "?sql_mode=STRICT_ALL_TABLES&tls=skip-verify&multiStatements=true&parseTime=true&charset=utf8mb4,utf8"
	client, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: client,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
