/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package commons

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"tuoxie-user-handle-service/config"
)

//默认数据库初始化
func init() {
	err := MySqlDBConfigInit()
	if err != nil {
		//@todo
		return
	}
}

//默认数据库初始化
func MySqlDBConfigInit() (err error) {
	data, err := json.Marshal(config.GetConfigData("mysql"))
	if err != nil {
		//@todo
		return
	}
	err = json.Unmarshal(data, MySqlDB)
	if err != nil {
		return
	}
	return
}
