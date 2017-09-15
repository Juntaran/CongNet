/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 1:25
 */

package controllers

import (
	"CongNet/models"
)


// register
type RegisterUserController struct {
	BaseController
}

func (this *RegisterUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/article", 302)
	} else {
		this.TplName = "register.tpl"
	}
}

func (this *RegisterUserController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")
	email := this.GetString("email")

	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	if "" == email {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写邮箱"}
		this.ServeJSON()
	}

	user := models.User{
		Name: 		name,
		Password: 	password,
		Email: 		email,
	}

	err := models.RegisterUser(user)

	if err == nil {
		//this.SetSession("userLogin", "1")
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	}
	this.ServeJSON()
	this.Redirect("/register", 302)
}


// login
type LoginUserController struct {
	BaseController
}

func (this *LoginUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/article", 302)
	} else {
		this.TplName = "login.tpl"
	}
}

func (this *LoginUserController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")

	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	err := models.LoginUser(name, password)

	if err == nil {
		this.SetSession("userLogin", "1")
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}


// logout
type LogoutUserController struct {
	BaseController
}

func (this *LogoutUserController) Get() {
	this.DelSession("userLogin")
	//this.Ctx.WriteString("you have logout")
	this.Redirect("/article", 302)
}