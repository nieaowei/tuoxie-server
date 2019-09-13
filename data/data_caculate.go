/*******************************************************
 *  ProjectName :tuoxie-server
 *  Author      :nieaowei
 *  Date        :2019/9/13
 *  Notes       :
 *******************************************************/
package data

import (
	"math"
	"net/http"
	"tuoxie-server/commons"
)

func DataCaculateHandler() {
	commons.MainRouter.HandleFunc("/getFall", getFallController)
}

func getFallController(w http.ResponseWriter, r *http.Request) {
	info := selectLastThreeaxisData(r.FormValue("username"))
	result := math.Sqrt(math.Pow(float64(info.X), 2) + math.Pow(float64(info.Y), 2) + math.Pow(float64(info.Z), 2))
	if result > 1.8 {

	}

}
