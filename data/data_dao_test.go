/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/7
 *  Notes       :
 *******************************************************/
package data

import (
	"reflect"
	"testing"
)

func Test_selectDataThree_axis(t *testing.T) {
	type args struct {
		tablename string
	}
	tests := []struct {
		name    string
		args    args
		wantRes *DataThree_axis
	}{
		// TODO: Add test cases.
		{
			"测试",
			args{tablename: "test"},
			&DataThree_axis{
				Time: "2019-07-15 01:00:00.0",
				X:    10,
				Y:    10,
				Z:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := selectDataThree_axis(tt.args.tablename); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("selectDataThree_axis() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

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
			"add user test.",
			args{username: string("test")},
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

func Test_addOneData_GPS(t *testing.T) {
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
			"add one data",
			args{
				username: "1nieaowei",
				data: DataThree_axis{
					Time: "",
					X:    1,
					Y:    10,
					Z:    20,
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
			"add one data",
			args{
				username: "1nieaowei",
				data: DataThree_axis{
					Time: "",
					X:    1,
					Y:    10,
					Z:    20,
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
