/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package data

import (
	"fmt"
	"log"
	"math"
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
		x1 := sensor.GetAccelVal(x)
		y1 := sensor.GetAccelVal(y)
		z1 := sensor.GetAccelVal(z)
		log.SetPrefix("Three")
		log.Println(x1, y1, z1)
		addOneData_Threeaxis(r.FormValue("username"), DataThree_axis{
			Time: "",
			X:    x1,
			Y:    y1,
			Z:    z1,
		})
		result := math.Sqrt(math.Pow(float64(x1), 2) + math.Pow(float64(y1), 2) + math.Pow(float64(z1), 2))
		log.SetPrefix("Three_result")
		log.Println(result)
		if result > 1.8 {
			temp := SelectLastGPSData(r.FormValue("username"))
			addOneData_Fall(r.FormValue("username"), temp)
		}
	}
}
