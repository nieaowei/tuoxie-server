/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"tuoxie-server/commons"
	"tuoxie-server/config"
	"tuoxie-server/data"
	"tuoxie-server/user"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("The User server is started."))
}

func main() {
	commons.MainRouter.HandleFunc("/", index)
	addr := flag.String("type", "user", "this is user server.")
	flag.Parse()
	switch *addr {
	case "user":
		{
			user.UserHandler()
			*addr = config.GetConfigData("server")["user_addr"].(string)
			fmt.Println("The user handle server is started.")
		}
	case "datasave":
		{
			data.DataSaveHandler()
			*addr = config.GetConfigData("server")["data_save_addr"].(string)
			fmt.Println("The data saving server is started.")
		}
	case "datahandle":
		{
			data.DataCaculateHandler()
			*addr = config.GetConfigData("server")["data_calculation"].(string)
			fmt.Println("The data calculation server is started.")
		}
	}

	http.ListenAndServe(*addr, commons.MainRouter)
}
