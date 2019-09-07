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
		longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 5)
		latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 5)
		fmt.Println(longitude, latitude)
		addOneData_GPS(r.FormValue("username"), DataGPS{
			Time:      "",
			Longitude: float32(longitude),
			Latitude:  float32(latitude),
		})
	case "data_three":
		x, _ := strconv.ParseFloat(r.FormValue("x"), 5)
		y, _ := strconv.ParseFloat(r.FormValue("y"), 5)
		z, _ := strconv.ParseFloat(r.FormValue("z"), 5)
		fmt.Println(x, y)
		addOneData_Threeaxis(r.FormValue("username"), DataThree_axis{
			Time: "",
			X:    float32(x),
			Y:    float32(y),
			Z:    float32(z),
		})
	}
}
