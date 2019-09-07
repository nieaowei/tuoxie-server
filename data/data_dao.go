package data

import (
	"fmt"
	"tuoxie-user-handle-service/commons"
)

//Read three-axis accelerometer data from the specified table
func selectData_Threeaxis(username string) (res *DataThree_axis) {
	fmt.Println(commons.TdDB)
	sqlStr := "select * from tuoxie.test"
	rows, err := commons.TdDB.Dql(sqlStr)
	fmt.Println(sqlStr)
	if err != nil {
		return
	}
	res = new(DataThree_axis)
	if rows.Next() {
		rows.Scan(&res.Time, &res.X, &res.Y, &res.Z)
		fmt.Println(res)
		return
	}
	return
}

func addOneData_GPS(username string, data DataThree_axis) (res int8) {
	sqlStr := "insert into ? values(?,?,?,?)"
	result, err := commons.TdDB.Dml(sqlStr, "gps_"+username, 0, data.X, data.Y, data.Z)
	if err != nil {
		return -1
	}
	if result == 0 {
		return -1
	}
	return
}

//
func addOneData_Threeaxis(username string, data DataThree_axis) (res int8) {
	sqlStr := "insert into ? values(?,?,?,?)"
	result, err := commons.TdDB.Dml(sqlStr, "threeaxis_"+username, 0, data.X, data.Y, data.Z)
	if err != nil {
		return -1
	}
	if result == 0 {
		return -1
	}
	return
}

//
func CreateTableToUser(username string) (res int8) {
	sqlStr := "create table threeaxis_" + username + "(uploadtime TIMESTAMP, x float, y float, z float)"
	fmt.Println(sqlStr)
	result, err := commons.TdDB.Dml(sqlStr)
	if err != nil {
		return -1
	}
	if result == 0 {
		return -1
	}
	sqlStr = "create table gps_" + username + "(uploadtime TIMESTAMP, x float, y float, z float)"
	result, err = commons.TdDB.Dml(sqlStr)
	if err != nil {
		return -1
	}
	if result == 0 {
		return -1
	}
	return
}

func selectLastGPSData(username string) (res *DataGPS) {
	sqlStr := "select last(*) from gps_" + username
	rows, err := commons.TdDB.Dql(sqlStr)
	if err != nil {
		return
	}
	if rows.Next() {
		res = new(DataGPS)
		rows.Scan(&res.Time, &res.longitude, &res.latitude)
	}
	return
}
