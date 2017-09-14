package main

import (
	_ "CongNet/routers"
	"github.com/astaxie/beego"
	"CongNet/models"
	"CongNet/controllers"
	"log"
)

func main() {
	// 从配置文件 conf/app.conf 中读取配置
	sqluser := beego.AppConfig.String("mysqluser")
	sqlpass := beego.AppConfig.String("mysqlpass")
	sqlurl  := beego.AppConfig.String("mysqlurls")
	sqldb   := beego.AppConfig.String("mysqldb")

	// 在 models 层创建数据库连接
	db, err := models.InitMySql(sqlurl, sqluser, sqlpass, sqldb)
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("db init success")
	}

	// 测试数据
	user := models.User{
		Name: "Juntaran",
		Password: "root",
		Email: "jacinthmail@gmail.com",
	}

	// 注册用户
	err = controllers.RegisterUser(db, user)
	if err != nil {
		log.Println(err)
		//panic(err)
	}

	// 登陆用户
	err = controllers.LoginUser(db, "Juntaran", "root")
	if err != nil {
		log.Println(err)
	}

	// 根据 Email 查找用户信息
	rets, _ := controllers.SearchUserByEmail(db, "jacinthmail@gmail.com")
	log.Println(rets)

	// 删除用户
	//err = controllers.DeleteUser(db, "Juntaran", "root")
	//if err != nil {
	//	log.Println(err)
	//}

	// 更改用户密码
	err = controllers.UpdateUserPassword(db, "Juntaran", "root", "jacinthmail@gmail.com", "root111root")
	if err != nil {
		log.Println(err)
	}

	// 根据 Email 查找用户信息
	rets, _ = controllers.SearchUserByEmail(db, "jacinthmail@gmail.com")
	log.Println(rets)

	// 更改用户邮箱
	err = controllers.UpdateUserEmail(db, "Juntaran", "root111root", "jacinthmail@sina.com")
	if err != nil {
		log.Println(err)
	}

	// 根据 Email 查找用户信息
	rets, _ = controllers.SearchUserByEmail(db, "jacinthmail@sina.com")
	log.Println(rets)

	//beego.Run()
}

