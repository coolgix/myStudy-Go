package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 如果当前路径没有MyDb.db，程序会自动创建
	dataSourceName := "root:123456@(127.0.0.1:3306)/test"
	db, _ := sql.Open("mysql", dataSourceName)

	// 通过程序执行SQL语句创建数据表
	sql_table := `CREATE TABLE IF NOT EXISTS userinfo (
	   id INT(11) PRIMARY KEY AUTO_INCREMENT,
	   username VARCHAR(64) NULL,
	   age INT(10) NULL,
	   created DATEtIME default CURRENT_TIMESTAMP
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	// 执行SQL语句
	db.Exec(sql_table)

	//新增数据
	stmt, _ := db.Prepare("INSERT INTO userinfo(username,age) values (?,?)")
	//传递参数并执行sql语句
	res, _ := stmt.Exec("Tom", "18")
	//返回新增数据ID
	id, _ := res.LastInsertId()
	fmt.Printf("新增数据的ID：%v\n", id)

	//批量新增数据
	UserList := [][]interface{}{{"Lily", 22}, {"Jim", 30}}
	for _, i := range UserList {
		//新增数据
		stmt, _ := db.Prepare("INSERT INTO userinfo(username, age) values(?, ?)")
		//传递参数并执行SQL语句
		res, _ := stmt.Exec(i[0], i[1])
		//返回新增数据的ID
		id, _ := res.LastInsertId()
		fmt.Printf("批量新增数据的ID：%v\n", id)
	}

	// 更新数据
	stmt, _ = db.Prepare("update userinfo set	username=? where id=?")
	//传递参数并执行SQL语句
	res, _ = stmt.Exec("Tim", 1)
	// 受影响的数据行数，返回int64类型数据
	affect, _ := res.RowsAffected()
	fmt.Printf("更新数据受影响的数据行数：%v\n", affect)

	//批量更新数据
	UserList1 := [][]interface{}{{"Betty", 3}, {"Jon", 4}}
	for _, i := range UserList1 {
		stmt, _ := db.Prepare("update userinfo set username=? where id=?")
		// 传递参数并执行SQL语句
		res, _ := stmt.Exec(i[0], i[1])
		// 受影响的数据行数，返回int64类型数据
		affect, _ := res.RowsAffected()
		fmt.Printf("更新数据受影响的数据行数：%v\n", affect)
	}
}
