package data

import (
	"fmt"
	"time"
	"tuoxie-user-handle-service/commons"
)

//Read three-axis accelerometer data from the specified table
func selectDataThree_axis(tablename string) (res *DataThree_axis) {
	fmt.Println(commons.TdDB)
	sqlStr := "select * from t"
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

//
func addDataThree_axis(tablename string, data *DataThree_axis) (res int8) {
	sqlStr := "insert into ? values(?,?,?,?)"
	timeStr := time.Now().String()
	result, err := commons.TdDB.Dml(sqlStr, timeStr, data.X, data.Y, data.Z)
	if err != nil {
		return -1
	}
	if result == 0 {
		return -1
	}
	return
}

func demo() {
}
