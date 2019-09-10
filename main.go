/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"net/http"
	"tuoxie-user-handle-service/commons"
	"tuoxie-user-handle-service/config"
	"tuoxie-user-handle-service/data"
	"tuoxie-user-handle-service/user"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("The User server is started."))
}

func main() {
	commons.MainRouter.HandleFunc("/", index)
	user.UserHandler()
	data.DataHandler()
	http.ListenAndServe(config.GetConfigData("server")["user_addr"].(string), commons.MainRouter)
}
