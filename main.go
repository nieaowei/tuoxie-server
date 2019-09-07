/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package main

import (
	"net/http"
	"tuoxie-user-handle-service/commons"
	"tuoxie-user-handle-service/config"
	"tuoxie-user-handle-service/user"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The User server is started."))
}

func main() {
	commons.MainRouter.HandleFunc("/", index)
	user.UserHandler()
	http.ListenAndServe(config.GetConfigData("server")["user_addr"].(string), commons.MainRouter)
}
