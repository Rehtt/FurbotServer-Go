package sql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) InitDB() (err error) {
	DB = &Database{}

	switch viper.GetString("db.use") {
	case "mysql":
		err = DB.initMySQL()
	case "sqlite":
		err = DB.initSQLite3()
	}
	if err != nil {
		return
	}
	dbset, _ := DB.Self.DB()
	dbset.SetMaxIdleConns(10) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。

	return nil
}

func (db *Database) initMySQL() (err error) {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.mysql.username"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.addr"),
		viper.GetInt("db.mysql.port"),
		viper.GetString("db.mysql.database"),
		true,
		"Local",
	)
	db.Self, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	return
}

func (db *Database) initSQLite3() (err error) {
	db.Self, err = gorm.Open(sqlite.Open(viper.GetString("db.sqlite.path")), &gorm.Config{})
	return
}
