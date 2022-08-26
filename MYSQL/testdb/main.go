package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// user:password@tcp(host:port)/database?charset-utf8mb4&loc=PRC&parseTime=true
	dsn := "root:zhanghaodemima19@tcp(127.0.0.1:3306)/zh?charset-utf8mb4" // 字符串的格式有对应的驱动

	db, err := sql.Open("mysql", dsn)

	fmt.Println(db, err)
	//测试能否连接
	fmt.Println(db.Ping())
}
