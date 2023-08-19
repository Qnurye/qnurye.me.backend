package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"qnurye/qnurye.me/pkg/config"
)

var db *gorm.DB

func init() {
	var err error
	cfg := config.Get()

	db, err = gorm.Open(getMysqlDialector(
		cfg.GetString("database.user"),
		cfg.GetString("database.password"),
		cfg.GetString("database.host"),
		cfg.GetInt("database.port"),
		cfg.GetString("database.name"),
	), &gorm.Config{})

	if err != nil {
		log.Fatal("Error on connecting to database:", err)
	}

}

func Get() *gorm.DB {
	return db
}

func getMysqlDialector(user string, pwd string, host string, port int, name string) gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, name))
}
