package main

import (
	_ "CongNet/routers"
	"github.com/astaxie/beego"
	//"CongNet/models"
	//"log"
)

func main() {

	//// 从配置文件 conf/app.conf 中读取配置
	//sqluser := beego.AppConfig.String("mysqluser")
	//sqlpass := beego.AppConfig.String("mysqlpass")
	//sqlurl  := beego.AppConfig.String("mysqlurls")
	//sqldb   := beego.AppConfig.String("mysqldb")
	//
	//// 在 models 层创建数据库连接
	//_, err := models.InitMySql(sqlurl, sqluser, sqlpass, sqldb)
	//if err != nil {
	//	log.Println(err)
	//	panic(err)
	//} else {
	//	log.Println("db init success")
	//}
	//
	//// 测试数据
	//user := models.User{
	//	Name: "Juntaran2",
	//	Password: "root2",
	//	Email: "jacinthmail2@gmail.com",
	//}
	//
	//// 注册用户
	//err := models.RegisterUser(user)
	//if err != nil {
	//	log.Println(err)
	//	//panic(err)
	//}
	//
	//// 登陆用户
	//err = models.LoginUser("Juntaran2", "root2")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//// 根据 Email 查找用户信息
	//rets, _ := models.SearchUserByEmail(db, "jacinthmail@gmail.com")
	//log.Println(rets)
	//
	// 删除用户
	//err = models.DeleteUser("Juntaran2", "rootroot","jacinthmail2@sina.com")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//// 更改用户密码
	//err = models.UpdateUserPassword(db, "Juntaran", "root", "jacinthmail@gmail.com", "root111root")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//// 根据 Email 查找用户信息
	//rets, _ = models.SearchUserByEmail(db, "jacinthmail@gmail.com")
	//log.Println(rets)
	//
	//// 更改用户邮箱
	//err = models.UpdateUserEmail(db, "Juntaran", "root111root", "jacinthmail@sina.com")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//// 根据 Email 查找用户信息
	//rets, _ = models.SearchUserByEmail(db, "jacinthmail@sina.com")
	//log.Println(rets)

	//models.StartAddFriend(3, 1, "111")
	//models.AcceptAddFriend(3, 1)
	//
	//models.StartAddFriend(3, 2, "222")
	//models.AcceptAddFriend(3, 2)

	//friends, length, _ := models.GetAllFriend(4)
	//log.Println(friends, length)

	//ret := models.GetCommentByDissID("0")
	//log.Println(ret)

	beego.Run()
}

