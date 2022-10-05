package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "golang"),
		beego.AppConfig.DefaultString("mysql::Password", "127.0.0.1"),
		beego.AppConfig.DefaultString("mysql::Host", "zhanghaodemima19"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
	)
	fmt.Println("dsn: ", dsn)
	db, _ = sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

}
