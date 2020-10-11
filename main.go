package main

import (
	"Data_Attestatio/db_mysql"
	_ "Data_Attestatio/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	//连接数据库
	db_mysql.DbOperation()
}

