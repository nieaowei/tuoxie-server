/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package data

//GPS data packet
type DataGPS struct {
	Time      string
	longitude float32
	latitude  float32
}

//Three-axis accelerometer data packet
type DataThree_axis struct {
	Time string
	X    float32
	Y    float32
	Z    float32
}
