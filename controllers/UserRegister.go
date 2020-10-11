package controllers

import (
	"Data_Attestatio/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//POST请求，用于处理前端注册页面请求
func (r *RegisterController) Post(){
	//解析前端数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("数据解析错误")
		return
	}
	fmt.Println("数据解析成功")
	//用户数据保存
	_, err = user.UserSave()
	if err != nil {
		r.Ctx.WriteString("用户保存失败")
		return
	}
	//保存成功，跳转登录
	r.TplName = "login.html"
}

