/*******************************************************
 *  ProjectName :tuoxie-user-handle-service
 *  Author      :nieaowei
 *  Date        :2019/9/7
 *  Notes       :
 *******************************************************/
package user

import (
	"reflect"
	"testing"
	"tuoxie-user-handle-service/commons"
)

func TestUser_LoginVerify(t *testing.T) {
	type fields struct {
		ID       int64
		Username string
		Password string
		Phone    string
		Created  string
		Updated  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes commons.Result
	}{
		// TODO: Add test cases.
		{
			"登录测试",
			fields{
				ID:       0,
				Username: "nieaowei",
				Password: "nieaowei",
				Phone:    "",
				Created:  "",
				Updated:  "",
			},
			commons.Result{
				Status: 200,
				Data:   nil,
				Msg:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Phone:    tt.fields.Phone,
				Created:  tt.fields.Created,
				Updated:  tt.fields.Updated,
			}
			if gotRes := u.LoginVerify(); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("LoginVerify() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
