package main

import (
	_ "CongNet/routers"
	"github.com/astaxie/beego"
	"CongNet/models"
	"fmt"
	"CongNet/models/login"
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
		fmt.Println(err)
	} else {
		fmt.Println("db success")
	}
	user := login.User{
		Name: "111",
		Password: "123",
		Email: "111",
	}

	err = db.CreateTable(login.User{}).Error
	if err != nil {
		panic(err)
	}

	// 主键为空返回`true`
	if db.NewRecord(user) == true {
		fmt.Println("主键为空")
	} else {
		fmt.Println("主键不为空")
	}

	db.Create(&user)

	if db.NewRecord(user) == true {
		fmt.Println("主键为空")
	} else {
		fmt.Println("主键不为空")
	}


	//beego.Run()
}

