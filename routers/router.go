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
	beego.Router("/user/:userid", &controllers.UserController{}, "*:RedID")	// 登陆后跳转到 /user/userid

	beego.Router("/friends", &controllers.FriendsGetController{})		// 测试用
	beego.Router("/diss", &controllers.CreateDissController{})			// 测试用
	beego.Router("/dissDel", &controllers.DeleteDissController{})		// 测试用
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