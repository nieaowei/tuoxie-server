/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config []byte
)

func init() {
	err := readConfig()
	if err != nil {
		//@todo
		panic(err)
	}
}

/*
Read configuration file.
*/
func readConfig() (err error) {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		//@todo
		data, err = ioutil.ReadFile("../config/config.json") //Used during testing.
		if err != nil {
			data, err = ioutil.ReadFile("./config.json") //Used during testing.
			if err != nil {
				return
			}
		}
	}
	config = data
	return
}

/*
Get the relevant configuration in the configuration file according to key.
*/
func GetConfigData(key string) (res map[string]interface{}) {
	temp := make(map[string]interface{})
	err := json.Unmarshal(config, &temp)
	if err != nil {
		//@todo
		return
	}
	return temp[key].(map[string]interface{})
}
