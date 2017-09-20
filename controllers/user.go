/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 1:25
 */

package controllers

import (
	"CongNet/models"
	"strconv"
	"log"
	"github.com/astaxie/beego"
	"fmt"
)


// register
type RegisterUserController struct {
	BaseController
}

func (this *RegisterUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/", 302)
	} else {
		this.TplName = "register.tpl"
	}
}

func (this *RegisterUserController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")
	email := this.GetString("email")

	if name == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	if email == "" {
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

type UserController struct {
	beego.Controller
}

func (this *UserController) RedID() {
	this.TplName = "index.tpl"
}

// login
type LoginUserController struct {
	BaseController
}

func (this *LoginUserController) Get() {
	if this.isLogin == true {
		// 1. 先根据 session 存储的 userEmail 查找 userid
		//log.Println("logincontroller:", this.userEmail)
		userid := models.GetIDByEmail(this.userEmail)
		log.Println("List userid:", userid)

		// 2. 跳转到 /user/userid
		//创建url
		//{{urlfor "UserController.List" ":name" "astaxie" ":age" "25"}}
		userController := UserController{}
		url := userController.URLFor("UserController.RedID", ":userid", strconv.Itoa(int(userid)))
		fmt.Println(url)
		log.Println("Redirect to", url)
		this.Redirect(url, 302)
	} else {
		this.TplName = "login.tpl"
	}
}

func (this *LoginUserController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")

	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写邮箱"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	isCancel, err := models.LoginUser(name, password)
	if err == nil {
		this.SetSession("userLogin", true)
		this.SetSession("userEmail", name)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		// 用户已经销号
		if isCancel == 1 {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "Farewell, this user is Canceled"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		}
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
	this.Redirect("/login", 302)
}


// cancel
type CancelUserController struct {
	BaseController
}

func (this *CancelUserController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "cancellation.tpl"
	} else {
		this.TplName = "login.tpl"
	}
}

func (this *CancelUserController) Post() {
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
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写注册邮箱"}
		this.ServeJSON()
	}

	err := models.DeleteUser(name, password, email)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Farewell~"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注销失败，请填写正确信息"}
	}
	this.DelSession("userLogin")
	this.ServeJSON()
}