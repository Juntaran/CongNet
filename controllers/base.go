/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/15 10:37 
  */

package controllers

import (
	"github.com/astaxie/beego"
)

// https://my.oschina.net/lockupme/blog/

type BaseController struct {
	beego.Controller
	isLogin		bool
	userEmail	string
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
		this.userEmail = ""
	} else {
		this.isLogin = true
		value, ok := this.GetSession("userEmail").(string)
		if !ok {
			this.isLogin = false
			this.userEmail = ""
		}
		this.userEmail = value
	}
	this.Data["isLogin"] = this.isLogin
}

func (this *BaseController) Go404() {
	this.TplName = "404.html"
	return
}
