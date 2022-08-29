package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func changeMoney(db *sql.Tx, id int, money float64) error {
	var accountMoney float64
	err := db.QueryRow("select money from account where id =?", id).Scan(&accountMoney)
	if err != nil {
		return err
	}

	if accountMoney < -money {
		return fmt.Errorf("现有金额为: %f, 金额不足", accountMoney)
	}
	_, err = db.Exec("update account set money=money+? where id=?", money, id)
	return err
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", "root", "zhanghaodemima19", "127.0.0.1", 3306, "zh")

	/*
			create table account(
				id bigint primary key auto_increment,
				name varchar(32) not null default '',
				money decimal(10,5) not null default 0
			)engine=innodb default charset utf8mb4;

		insert into account(name, money) values("bob", 1000);
		insert into account(name, money) values("mike", 1000);
	*/

	var a, b int = 1, 2

	db, _ := sql.Open("mysql", dsn)
	//同时失败同事成功
	//事物
	tx, _ := db.Begin()
	m := 500.
	err1 := changeMoney(tx, a, -m)
	err2 := changeMoney(tx, b, m)

	if err1 != nil && err2 != nil {
		tx.Commit()
	}

	fmt.Println(err1, err2)
}
