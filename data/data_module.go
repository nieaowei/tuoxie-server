/*******************************************************
 *  ProjectName :tuoxie-data-save-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package data

//GPS data packet
type DataGPS struct {
}

//Three-axis accelerometer data packet
type DataThree_axis struct {
	Time string
	X    int
	y    int
	Z    int
}
