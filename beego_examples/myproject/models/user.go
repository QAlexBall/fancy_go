package models

import (
	"fmt"
	"myproject/utils"
)

// User model
type User struct {
	ID         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

// InsertUser =>
func InsertUser(user User) (int64, error) {
	sql := fmt.Sprintf("insert into beego_users values (%s, %s, %d, %d)",
		user.Username, user.Password, user.Status, user.Createtime)
	fmt.Println(sql)
	return utils.ModifyDB(sql)
}

// QueryUserWightCon =>
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// QueryUserWithUsername =>
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}
