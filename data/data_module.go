/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package data

//GPS data packet
type DataGPS struct {
	Time      string  `json:"time"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

//Three-axis accelerometer data packet
type DataThree_axis struct {
	Time string  `json:"time"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}
