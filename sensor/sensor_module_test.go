/*******************************************************
 *  ProjectName :tuoxie-server
 *  Author      :nieaowei
 *  Date        :2019/9/13
 *  Notes       :
 *******************************************************/
package sensor

import "testing"

func TestGetAccelVal(t *testing.T) {
	type args struct {
		id int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes float32
	}{
		// TODO: Add test cases.
		{
			"正数测试",
			args{id: 32},
			-1,
		},
		{
			"负数测试",
			args{id: -32},
			-1.500,
		},
		{
			"0测试",
			args{id: 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetAccelVal(tt.args.id); gotRes != tt.wantRes {
				t.Errorf("GetAccelVal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
