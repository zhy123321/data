package controllers

import (
	"Data_Attestatio/models"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}

//前端直接登录请求
func (l *LoginControllers) Get(){
	l.TplName = "login.html"
}

//登录请求
func (l *LoginControllers) Post(){
	//解析前端登录信息
	var user models.User
	err := l.ParseForm(&user)
	if err != nil{
		l.Ctx.WriteString("数据解析错误")
	}
	//用户数据查询
	u, err := user.UserData()
	if err != nil{
		l.Ctx.WriteString("用户登录失败")
		return
	}
	l.Data["Phone"] = u.Phone
	l.TplName = "corepage.html"
}
