package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:c)t9e5wku&i9@tcp(cs-cn-east-73.teamcode.com:4018)/test?charset=utf8")
	checkErr(err)
	db.Ping()
	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
