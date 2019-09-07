/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package config

import (
	"reflect"
	"testing"
)

func TestGetConfigData(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]interface{}
	}{
		// TODO: Add test cases.
		{
			"获取数据",
			args{key: "server"},
			map[string]interface{}{
				"data_calculation": "localhost:8080",
				"data_save_addr":   "localhost:8080",
				"user_addr":        "localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetConfigData(tt.args.key); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetConfigData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
