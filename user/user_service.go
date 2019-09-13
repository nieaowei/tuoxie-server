/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package user

import (
	"regexp"
	"tuoxie-server/commons"
	"tuoxie-server/data"
)

/*
	To verify that the user's account and password match.
*/
func (u *User) LoginVerify() (res commons.Result) {
	res.Status = 200
	switch {
	case u.Username != "":
		user := selectByPwd(u.Username, u.Password)
		if user != nil {
			return
		}
	case u.Phone != "":
		user := selectByPwd(u.Username, u.Password)
		if user != nil {
			return
		}
	}
	res.Status = 400
	return
}

/*
	Format matching the mailbox and mobile phone number ,and call the database opration to add a user.
*/
func (u *User) RegisterVerify() (res commons.Result) {
	if !u.MatchPhone() {
		res.Status = DATA_PHONE_ERROR
		res.Msg = "手机号码格式不正确"
		return
	}
	if res1 := addUserByUPP(u); res1 != USEROK {
		res.Status = int(res1)
		switch res1 {
		case EXIST_USERNAME:
			res.Msg = "用户名已经存在"
		case EXIST_PHONE:
			res.Msg = "手机号码已经存在"
		}
		return
	}
	result := data.CreateTableToUser(u.Username)
	if result != 0 {
		res.Status = 400
		res.Msg = "no known error"
		return
	}
	res.Status = 200
	return
}

/*
	Use regular expressions to match whether the data satisfies the format.
*/
func (u *User) MatchPhone() (res bool) {
	res, _ = regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9][\d]{8})$`, u.Phone)
	return
}

/*
	get last location
*/
func (u *User) GetLastGPS() (res commons.Result) {
	gpsinfo := data.SelectLastGPSData(u.Username)
	res.Status = 400
	if gpsinfo != nil {
		res.Status = 200
		res.Data = gpsinfo
	}
	return
}

/*

 */

func (u *User) GetLastFall() (res commons.Result) {
	gpsinfo := data.SelectLastFallData(u.Username)
	res.Status = 400
	if gpsinfo != nil {
		res.Status = 200
		res.Data = gpsinfo
	}
	return
}
