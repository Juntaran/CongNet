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
	beego.Router("/friends", &controllers.FriendsGetController{})
	//beego.Router("/user/:id", &controllers.FriendsGetController{})

	//beego.Router("/user/", &controllers.LoginUserController{})
	beego.Router("/user/:userid", &controllers.LoginUserController{}, "*:List")
}

/*
//注册路由
beego.Router("/user/list/:name/:age", &userController, "*:List")

//创建url
//{{urlfor "UserController.List" ":name" "astaxie" ":age" "25"}}

url := userController.UrlFor("UserController.List", ":name", "astaxie", ":age", "25")

//输出 /user/list/astaxie/25
fmt.Println(url)
*/