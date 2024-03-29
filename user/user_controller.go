/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tuoxie-server/commons"
)

func UserHandler() {
	commons.MainRouter.HandleFunc("/login", loginController)
	commons.MainRouter.HandleFunc("/register", registerController)
	commons.MainRouter.HandleFunc("/getLastGPS", getLastGPSController)
	commons.MainRouter.HandleFunc("/getLastFall", getLastFallController)
}

func getLastFallController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fmt.Println("[getGPS]:" + username)
	user := User{
		ID:       0,
		Username: username,
		Password: "",
		Phone:    "",
		Created:  "",
		Updated:  "",
	}
	res := user.GetLastFall()
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)

}

func getLastGPSController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fmt.Println("[getGPS]:" + username)
	user := User{
		ID:       0,
		Username: username,
		Password: "",
		Phone:    "",
		Created:  "",
		Updated:  "",
	}
	res := user.GetLastGPS()
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}

func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	user := User{
		ID:       0,
		Username: username,
		Password: password,
		Phone:    username,
		Created:  "",
		Updated:  "",
	}
	res := user.LoginVerify()
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}

func registerController(w http.ResponseWriter, r *http.Request) {
	user := NewUserByRequest(r) //不知道效率
	res := user.RegisterVerify()
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}
