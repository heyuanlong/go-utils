package mysql

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	kconf "github.com/heyuanlong/go-utils/common/conf"
	klog "github.com/heyuanlong/go-utils/common/log"
)

var (
	MysqlClient     *sql.DB
)

func init() {
	
}

func InitMysql()  {

	user,_ := 		kconf.GetString("mysql","user")
	password,_ := 	kconf.GetString("mysql","password")
	ip,_ := 		kconf.GetString("mysql","ip")
	port,_ := 		kconf.GetString("mysql","port")
	mysqldb,_ := 	kconf.GetString("mysql","mysqldb")

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		user,
		password,
		ip,
		port,
		mysqldb,
	)
	MysqlClient, _ = sql.Open("mysql", addr)
	MysqlClient.SetMaxOpenConns(2000)
	MysqlClient.SetMaxIdleConns(10)
	MysqlClient.Ping()
	klog.Warn.Printf("mysql open ok")
}


