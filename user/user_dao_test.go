/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/6
 *  Notes       :
 *******************************************************/
package user

import "testing"

func Test_addUserByUPP(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{
			"增加用户测试",
			args{&User{
				ID:       1,
				Username: "nieaowei",
				Password: "nieaowei",
				Phone:    "13575132575",
				Created:  "",
				Updated:  "",
			}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addUserByUPP(tt.args.u); got != tt.want {
				t.Errorf("addUserByUPP() = %v, want %v", got, tt.want)
			}
		})
	}
}
