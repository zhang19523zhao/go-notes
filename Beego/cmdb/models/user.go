package models

import (
	"cmdb/utils"
	"fmt"
	"time"
)

const (
	sqlQueryByName = `select id, name, password from user where name=?`
	sqlQuery       = `select id, staff_id, name,nickname,gender,tel,email,addr,department,status,created_at,updated_at,deleted_at from user`
	sqlQueryName   = `select id, staff_id, name,nickname,gender,tel,email,addr,department,status,created_at,updated_at,deleted_at from user wherename=?`
)

// User 用户对象
type User struct {
	ID         int
	StaffID    string
	Name       string
	Nickname   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
	Created_at *time.Time
	Updated_at *time.Time
	Deleted_at *time.Time
}

// ValidPassword 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return utils.Md5Text(password) == u.Password
}

// 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{}
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		fmt.Println("gertuser: ", user)
		return user
	} else {
		fmt.Println("getuser err: ", err)
	}
	return nil
}

// QureyUser 查询用户
func QueryUser(q string) []*User {
	users := make([]*User, 0, 10)
	sqlq := sqlQuery

	q = utils.Like(q)
	parms := []interface{}{}
	if q != "" {
		sqlq += " WHERE staff_id like  ? ESCAPE '/' or name like  ? ESCAPE '/' or nickname like ? ESCAPE '/'  or email like ? ESCAPE '/'  or tel like ? ESCAPE '/'   or department like ? ESCAPE '/' or tel like ? ESCAPE '/'   or addr like ? ESCAPE '/' "
		//fmt.Println(sqlq)
		//fmt.Println("q=", q)
		parms = append(parms, q, q, q, q, q, q, q, q)
	}

	rows, err := db.Query(sqlq, parms...)
	//fmt.Println(err)
	if err != nil {
		return users
	}
	for rows.Next() {
		user := &User{}
		//id, staff_id, name,nickname,gender,tel,email,addr,department,status,created_at,updated_at,deleted
		if err := rows.Scan(&user.ID, &user.StaffID, &user.Name, &user.Nickname, &user.Gender, &user.Tel, &user.Email, &user.Addr, &user.Department, &user.Status, &user.Created_at, &user.Updated_at, &user.Deleted_at); err == nil {
			users = append(users, user)
		} else {
			fmt.Println(err)
		}

	}
	return users
}

// 显示性别
func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}

// 显示状态
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}
