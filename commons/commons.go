/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package commons

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
)

//客户端服务端交互模板
type Result struct {
	Status int         `json:"status"` //状态为200成功
	Data   interface{} `json:"data"`   //返回数据
	Msg    string      `json:"msg"`    //返回信息
}

var (
	MainRouter = mux.NewRouter()
	MySqlDB    = new(MyDataBase)
	TdDB       = new(MyDataBase)
)

type MyDataBase struct {
	dB         *sql.DB
	stmt       *sql.Stmt
	rows       *sql.Rows
	Driver     string `json:"driver"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Table      string `json:"table"`
	DataSource string `json:"dataSource"`
}

//打开数据库链接
func (this *MyDataBase) openConn() (err error) {
	this.dB, err = sql.Open(this.Driver, this.Name+":"+this.Password+"@tcp("+this.DataSource+")/"+this.Table)
	if err != nil {
		//@TODO 实现日志文件记录
		fmt.Println("open database fail.")
		return
	}
	return nil
}

//关闭数据库链接
func (this *MyDataBase) CloseConn() {
	if this.rows != nil {
		this.rows.Close()
	}
	if this.stmt != nil {
		this.stmt.Close()
	}
	if this.dB != nil {
		this.dB.Close()
	}

}

//
func (this *MyDataBase) Dml(sql string, args ...interface{}) (lenth int64, err error) {
	err = this.openConn()
	defer this.CloseConn()
	if err != nil {
		//@TODO 日志
		fmt.Println("open database fail")
		return
	}
	this.stmt, err = this.dB.Prepare(sql)
	if err != nil {
		//@TODO
		fmt.Println("prepare fail")
		return
	}
	result, err := this.stmt.Exec(args...)
	if err != nil {
		//@TODO
		fmt.Println("exec fail")
		return
	}
	lenth, err = result.RowsAffected()
	if err != nil {
		//@todo
		fmt.Println("no res")
		return
	}
	return
}

func (this *MyDataBase) Dql(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	err = this.openConn()
	if err != nil {
		//@TODO 日志
		fmt.Println("打开数据库失败")
		return
	}
	this.stmt, err = this.dB.Prepare(sql)
	if err != nil {
		//@TODO
		fmt.Println("预处理失败")
		return
	}
	rows, err = this.stmt.Query(args...)
	if err != nil {
		//@TODO
		fmt.Println("执行失败")
		return
	}
	//this.CloseConn()//调用此函数要记得关闭
	return
}

//if you alter database, you use this func.
func (this *MyDataBase) SetTable(name string) {
	this.Table = name
}
