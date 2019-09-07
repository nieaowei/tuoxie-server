/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package user

import (
	"fmt"
	"regexp"
	"time"
	"tuoxie-user-handle-service/commons"
)

/*
Query data in the database by account number and password .
Return nil if the query has an error or the data does not exist.
*/
func selectByPwd(usern, pwd string) (user *User) {
	sql := "select * from user where username=? and password=? or phone=? and password=?"
	rows, err := commons.MySqlDB.Dql(sql, usern, pwd, usern, pwd)
	if err != nil {
		//@todo
		return
	}
	if rows.Next() {
		user = new(User)
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Created, &user.Updated)
		if err != nil {
			//@todo 查询出错
			return nil
		}
		commons.MySqlDB.CloseConn()
		return
	}
	return
}

/*
Add a record to the database based on the user object and be able to analyze the error.
According to the return value , you can know the cause of the error and can handle it
according to different reasons
*/
func addUserByUPP(u *User) int8 {
	/*
		下列操作效率低
	*/
	//sql := "select * from tb_user where username=?"
	//rows , err := commons.MySqlDB.Dql(sql,username)
	//if err != nil {
	//	//@todo
	//	commons.MySqlDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MySqlDB.CloseConn()
	//	return EXIST_USERNAME
	//}
	//sql = "select * from tb_user where phone=?"
	//rows , err = commons.MySqlDB.Dql(sql,phone)
	//if err !=nil{
	//	commons.MySqlDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MySqlDB.CloseConn()
	//	return EXIST_PHONE
	//}
	//sql = "select * from tb_user where email=?"
	//rows , err = commons.MySqlDB.Dql(sql,email)
	//if err !=nil{
	//	commons.MySqlDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MySqlDB.CloseConn()
	//	return EXIST_EMAIL
	//}
	//If result is existed, the following code will be not run.

	// The following code implements the function of error location bt analyzing the error information.
	//And improve operational efficiency.
	sql := "insert into user values(default,?,?,?,?,?)"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	lenth, err := commons.MySqlDB.Dml(sql, u.Username, u.Password, u.Phone, timeStr, timeStr)
	if err != nil { //解析错误
		r, _ := regexp.Compile("'[a-zA-Z0-9]+'")
		temp := r.FindAllString(fmt.Sprintf("%s", err), 2)
		switch temp[len(temp)-1] {
		// There are a point in here,you should not judge temp[2] directly,otherwise result in panic.
		// The solution is to take the judgment by acquiring the length of the data slice.
		case "'username'":
			return EXIST_USERNAME
		case "'phone'":
			return EXIST_PHONE
		case "'email'":
			return EXIST_EMAIL
		default:
			return USERERROR
		}
	}
	if lenth < 0 {
		return USERERROR
	}
	return USEROK
}
