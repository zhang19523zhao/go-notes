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
	//fmt.Println(db.Ping())

	//执行
	//sql => go 字符串
	//fmt.Println(db.Exec(`
	//create table if not  exists  testdb(
	//    id bigint primary key auto_increment,
	//    name varchar(15) not null default '' comment 'testdb名字'
	//)engine=innodb default charset utf8mb4;
	//`))

	//status := fmt.Sprintf("insert into testdb values(%d, '%s');", 1111, "bob2")
	//status2 := "insert into testdb values(?, ?)"
	//result2, _ := db.Exec(status2, 1113, "resutl2")
	//fmt.Println(result2.RowsAffected())
	//fmt.Println(status)
	//result, _ := db.Exec(`insert into testdb values(1411, "bob")`)
	////result, _ = db.Exec(`insert into testdb values(11, "bob")`)
	//result, _ = db.Exec(status)
	//fmt.Println(result.RowsAffected())
	//fmt.Println(result.LastInsertId())
	//读
	//rows, _ := db.Query("select id, name from testdb;")

	var (
		id   int
		name string
	)
	//for rows.Next() {
	//	rows.Scan(&id, &name)
	//	fmt.Println(id, name)
	//}

	//直接查询一行
	rows := db.QueryRow("select id, name from testdb where id > ?", 10)
	rows.Scan(&id, &name)
	fmt.Println(id, name)
	fmt.Println(rows.Scan(id, name))
}
