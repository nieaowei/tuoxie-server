package data

import (
	"reflect"
	"testing"
)

func TestCreateTableToUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int8
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{username: "test3"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := CreateTableToUser(tt.args.username); gotRes != tt.wantRes {
				t.Errorf("CreateTableToUser() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_addOneData_Threeaxis(t *testing.T) {
	type args struct {
		username string
		data     DataThree_axis
	}
	tests := []struct {
		name    string
		args    args
		wantRes int8
	}{
		// TODO: Add test cases.
		{
			"test3",
			args{
				username: "test3",
				data: DataThree_axis{
					"",
					10,
					12,
					13,
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := addOneData_Threeaxis(tt.args.username, tt.args.data); gotRes != tt.wantRes {
				t.Errorf("addOneData_Threeaxis() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_selectData_Threeaxis(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantRes *DataThree_axis
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := selectData_Threeaxis(tt.args.username); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("selectData_Threeaxis() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_selectLastGPSData(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantRes *DataGPS
	}{
		// TODO: Add test cases.
		{
			"test",
			args{username: "test3"},
			&DataGPS{
				Time:      "2019-09-07 23:02:51.718",
				Longitude: 10,
				Latitude:  32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := SelectLastGPSData(tt.args.username); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SelectLastGPSData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_addOneData_GPS(t *testing.T) {
	type args struct {
		username string
		data     DataGPS
	}
	tests := []struct {
		name    string
		args    args
		wantRes int8
	}{
		// TODO: Add test cases.
		{
			"test",
			args{
				username: "test3",
				data: DataGPS{
					Time:      "",
					Longitude: 10,
					Latitude:  32,
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := addOneData_GPS(tt.args.username, tt.args.data); gotRes != tt.wantRes {
				t.Errorf("addOneData_GPS() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_selectLastThreeaxisData(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantRes *DataThree_axis
	}{
		// TODO: Add test cases.
		{
			"test",
			args{username: "test3"},
			&DataThree_axis{
				Time: "2019-09-07 23:04:01.911",
				X:    10,
				Y:    12,
				Z:    13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := selectLastThreeaxisData(tt.args.username); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("selectLastThreeaxisData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
