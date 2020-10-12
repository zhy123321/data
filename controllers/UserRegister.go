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
		fmt.Println(err.Error())
		return
		r.Ctx.WriteString("数据解析错误")
	}
	fmt.Println("数据解析成功")
	//用户数据保存
	id, err := user.UserSave()
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)
	//保存成功，跳转登录
	r.TplName = "login.html"
}

