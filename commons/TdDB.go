/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package commons

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
	"tuoxie-user-handle-service/config"
	_ "tuoxie-user-handle-service/taosSql"
)

//默认数据库初始化
func init() {
	err := TdDBConfigInit()
	if err != nil {
		//@todo
		return
	}
}

//默认数据库初始化
func TdDBConfigInit() (err error) {
	data, err := json.Marshal(config.GetConfigData("tdengine"))
	if err != nil {
		//@todo
		return
	}
	err = json.Unmarshal(data, TdDB)
	if err != nil {
		return
	}
	return
}

//创建数据库
func create_database_stmt(db *sql.DB, demodb string) {
	st := time.Now().Nanosecond()
	// create database
	//var stmt interface{}
	stmt, err := db.Prepare("create database ?")
	checkErr(err)

	//var res driver.Result
	res, err := stmt.Exec(demodb)
	checkErr(err)

	//fmt.Printf("Query OK, %d row(s) affected()", res.RowsAffected())
	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("create database result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1e9)
}

//使用数据库
func use_database_stmt(db *sql.DB, demodb string) {
	st := time.Now().Nanosecond()
	// create database
	//var stmt interface{}
	stmt, err := db.Prepare("use " + demodb)
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("use database result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1e9)
}

//创建数据表
func create_table_stmt(db *sql.DB, demot string) {
	st := time.Now().Nanosecond()
	// create table
	// (ts timestamp, id int, name binary(8), len tinyint, flag bool, notes binary(8), fv float, dv double)
	stmt, err := db.Prepare("create table ? (? timestamp, ? int, ? binary(8), ? tinyint, ? bool, ? binary(8), ? float, ? double)")
	checkErr(err)

	res, err := stmt.Exec(demot, "ts", "id", "name", "len", "flag", "notes", "fv", "dv")
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("create table result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1e9)
}

//插入数据
func insert_data_stmt(db *sql.DB, demot string) {
	st := time.Now().Nanosecond()
	// insert data into table
	stmt, err := db.Prepare("insert into ? values(?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(demot, "now", 1000, "'haidian'", 6, true, "'AI world'", 6987.654, 321.987,
		"now+1s", 1001, "'changyang'", 7, false, "'DeepMode'", 12356.456, 128634.456,
		"now+2s", 1002, "'chuangping'", 8, true, "'database'", 3879.456, 65433478.456)
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1e9)
}

//查询数据
func select_data_stmt(db *sql.DB, demot string) {
	st := time.Now().Nanosecond()

	stmt, err := db.Prepare("select ?, ?, ?, ?, ?, ?, ?, ? from ?") // go binary mode
	checkErr(err)

	rows, err := stmt.Query("ts", "id", "name", "len", "flag", "notes", "fv", "dv", demot)
	checkErr(err)

	fmt.Printf("%10s%s%8s %5s %9s%s %s %8s%s %7s%s %8s%s %11s%s %14s%s\n", " ", "ts", " ", "id", " ", "name", " ", "len", " ", "flag", " ", "notes", " ", "fv", " ", " ", "dv")
	var affectd int
	for rows.Next() {
		var ts string
		var name string
		var id int
		var len int8
		var flag bool
		var notes string
		var fv float32
		var dv float64

		err = rows.Scan(&ts, &id, &name, &len, &flag, &notes, &fv, &dv)
		//fmt.Println("start scan fields from row.rs, &fv:", &fv)
		//err = rows.Scan(&fv)
		checkErr(err)

		fmt.Printf("%s\t", ts)
		fmt.Printf("%d\t", id)
		fmt.Printf("%10s\t", name)
		fmt.Printf("%d\t", len)
		fmt.Printf("%t\t", flag)
		fmt.Printf("%s\t", notes)
		fmt.Printf("%06.3f\t", fv)
		fmt.Printf("%09.6f\n", dv)

		affectd++

	}

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1e9)
}

//检测错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
