package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	kconf "github.com/heyuanlong/go-utils/common/conf"
	klog "github.com/heyuanlong/go-utils/common/log"
)

var (
	GormDB     *gorm.DB
)

func InitGorm()  {

	user,_ := 		kconf.GetString("mysql","user")
	password,_ := 	kconf.GetString("mysql","password")
	ip,_ := 		kconf.GetString("mysql","ip")
	port,_ := 		kconf.GetString("mysql","port")
	mysqldb,_ := 	kconf.GetString("mysql","mysqldb")

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		ip,
		port,
		mysqldb,
	)
	var err error
	GormDB, err = gorm.Open("mysql", addr)
	if err == nil {
		klog.Warn.Println("gorm open ok")
	}else{
		klog.Warn.Println("gorm open fail:",err)
	}
}

