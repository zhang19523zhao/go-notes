package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//字符串的格式由对应的字符串驱动定义
	//"user:passwd@tcp(host:port)?charset=utf8mb4&loc=PRC"
	dsn := "root:zhanghaodemima19@(localhost:3306)?charset=utf8mb4&loc=PRC"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	//执行
	fmt.Println(db.Exec(`
	create table if not exists testzh(
	    id biging primary key auto_increment,
	    name varchar(32) not null default '' comment 'testzh名字'
	)engine=innodb default charset utf8mb4;
`))
}
