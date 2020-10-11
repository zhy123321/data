package controllers

import "github.com/astaxie/beego"

type LoginControllers struct {
	beego.Controller
}

func (l *LoginControllers) Get(){
	l.TplName = "login.html"
}