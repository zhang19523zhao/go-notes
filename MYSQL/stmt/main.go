package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", "root", "zhanghaodemima19", "127.0.0.1", 3306, "zh")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset utf8mb4", "root", "zhanghaodemima19", "127.0.0.1", 3306, "zh")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	//_, err = db.Exec("insert into account values(?,?,?)", 0, "BOB", 0)
	fmt.Println(err)
	stmt, _ := db.Prepare("insert into account(name, money)  values(?,?)")

	start := time.Now()
	for i := 0; i <= 100000; i++ {
		//_, err := db.Exec("insert into account(name, money)  values(?,?)", "BOB", i)
		stmt.Exec("bob", i)
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println(time.Now().Sub(start))
}
