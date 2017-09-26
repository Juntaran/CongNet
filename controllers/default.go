package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplName = "index.html"
	//c.Ctx.WriteString("hello")
	//url := this.URLFor("LoginUserController.Get", ":page", "111")
	//log.Println("default url:", url)
	//this.Redirect(url, 302)
	this.Redirect("/login", 302)
}