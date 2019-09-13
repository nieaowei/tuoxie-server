/*******************************************************
 *  ProjectName :tuoxie-server
 *  Author      :nieaowei
 *  Date        :2019/9/13
 *  Notes       :
 *******************************************************/
package sensor

var sensorArray = []float32{
	0, 0.047, 0.094, 0.141, 0.188, 0.234, 0.281, 0.328, 0.375, 0.422, 0.469, 0.516, 0.563, 0.609, 0.656,
	0.703, 0.750, 0.797, 0.844, 0.891, 0.938, 0.984, 1.031, 1.078, 1.125, 1.172, 1.219, 1.266, 1.313, 1.359, 1.406, 1.453, 1.500,
}

func GetAccelVal(id int) (res float32) {
	if id >= 32 || id < -32 {
		return -1
	}
	if id < 0 {
		return -sensorArray[-id]
	}
	return sensorArray[id]
}
