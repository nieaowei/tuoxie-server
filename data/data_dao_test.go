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
			args{tablename: "t"},
			&DataThree_axis{
				Time: "",
				X:    0,
				Y:    0,
				Z:    0,
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
