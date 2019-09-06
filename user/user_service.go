/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package user

import (
	"regexp"
	"tuoxie-user-handle-service/commons"
)

/*
	This method will be not used.
*/
func loginService(usern, pwd string) (res commons.Result) {
	user := selectByPwd(usern, pwd)
	if user != nil { //查询到数据
		res.Status = 200
	} else { //没有数据
		res.Status = 400
	}
	return
}

//func registerService(username, pwd, phone,email string) (res commons.Result) {
//	//首先验证是否符合要求
//	if !matchPhone(phone) {
//		res.Status = 200
//		return
//	}
//	if !matchEmail(email) {
//		res.Status = 400
//		return
//	}
//
//	//如果都符合要求继续执行注册
//	err := addUserByUPP(username, pwd, phone,email)
//	if err != 0 {
//		res.Status = 400
//		return
//	} else {
//		res.Status = 200
//		return
//	}
//}

func matchPhone(phone string) (res bool) {
	res, _ = regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9][\d]{8})$`, phone)
	return
}

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
