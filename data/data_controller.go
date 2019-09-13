/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package data

import (
	"fmt"
	"net/http"
	"strconv"
	"tuoxie-server/commons"
	"tuoxie-server/sensor"
)

func DataSaveHandler() {
	commons.MainRouter.HandleFunc("/postinfo", postInfoController)
}

func postInfoController(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("type") {
	case "data_gps":
		longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 64)
		latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 64)
		if longitude == 0 && latitude == 0 {
			return
		}
		fmt.Println(longitude, latitude)
		addOneData_GPS(r.FormValue("username"), DataGPS{
			Time:      "",
			Longitude: longitude / 100,
			Latitude:  latitude / 100,
		})
	case "data_three":
		x, _ := strconv.Atoi(r.FormValue("x"))
		y, _ := strconv.Atoi(r.FormValue("y"))
		z, _ := strconv.Atoi(r.FormValue("z"))
		fmt.Println(x, y, z)
		addOneData_Threeaxis(r.FormValue("username"), DataThree_axis{
			Time: "",
			X:    sensor.GetAccelVal(x),
			Y:    sensor.GetAccelVal(y),
			Z:    sensor.GetAccelVal(z),
		})
	}
}
