package sqldemo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"

	"util"
)

func dsn() string {
	return "godemo:godemo123@tcp(127.0.0.1:3306)/godemo?charset=utf8"
}

func MysqlTest() {
	//打开数据库mytest
	fmt.Println("open the database, godemo")

	db, err := sql.Open("mysql", dsn())

	util.CheckErr(err)

	//插入数据
	fmt.Println("insert data info into table, userinfo")

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")

	util.CheckErr(err)

	res, err := stmt.Exec("Test", "people", "2017-10-27")

	util.CheckErr(err)

	//记录ID
	fmt.Println("record the ID of affecting item")

	id, err := res.LastInsertId()

	util.CheckErr(err)

	fmt.Println(id)

	//更新数据
	fmt.Println("update data info where id was conform to be affected")

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")

	util.CheckErr(err)

	res, err = stmt.Exec("man", id)

	util.CheckErr(err)

	affect, err := res.RowsAffected()

	util.CheckErr(err)

	fmt.Println(affect)

	//查询数据
	fmt.Println("query data info from table, userinfo")

	rows, err := db.Query("SELECT * FROM userinfo")

	util.CheckErr(err)

	fmt.Println("iterator data info to show")

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)

		util.CheckErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//
	//util.CheckErr(err)
	//
	//res, err = stmt.Exec(id)
	//
	//util.CheckErr(err)
	//
	//affect, err = res.RowsAffected()
	//
	//util.CheckErr(err)
	//
	//fmt.Println(affect)

	//关闭数据库mytest
	db.Close()
}
