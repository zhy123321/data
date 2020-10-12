package main

import (
	"Data_Attestatio/db_mysql"
	_ "Data_Attestatio/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库
	db_mysql.DbOperation()
	beego.Run()
}

