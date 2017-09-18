package routers

import (
	"CongNet/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginUserController{})
	beego.Router("/logout", &controllers.LogoutUserController{})
	beego.Router("/register", &controllers.RegisterUserController{})
	beego.Router("/cancellation", &controllers.CancelUserController{})
	beego.Router("/friends", &controllers.GetFriendsController{})
}
