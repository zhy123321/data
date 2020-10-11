package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func DbOperation()  {

	//连接数据库
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser +":"+ dbPassword + "@tcp(" + dbIP + ")/" + dbName + "?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接错误")
	}
	fmt.Println("连接数据库成功")
	Db = db
}


