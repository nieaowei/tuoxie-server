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
	"tuoxie-user-handle-service/commons"
)

func DataHandler() {
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
		x, _ := strconv.ParseFloat(r.FormValue("x"), 32)
		y, _ := strconv.ParseFloat(r.FormValue("y"), 32)
		z, _ := strconv.ParseFloat(r.FormValue("z"), 32)
		fmt.Println(x, y, z)
		addOneData_Threeaxis(r.FormValue("username"), DataThree_axis{
			Time: "",
			X:    float32(x),
			Y:    float32(y),
			Z:    float32(z),
		})
	}
}
