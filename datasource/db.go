package datasource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iris-gorm-demo/conf"
	"strings"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	path := strings.Join([]string{conf.Sysconfig.DBUserName, ":", conf.Sysconfig.DBPassword,
		"@(", conf.Sysconfig.DBIp, ":", conf.Sysconfig.DBPort, ")/",
		conf.Sysconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	println(db)
}
