/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/14 18:03 
  */

package models

import (
	"github.com/astaxie/beego"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"database/sql"

	"log"
	"fmt"
)

// mysql 数据库初始化
// 1. 连接数据库
// 2. 创建必须的数据库表

var conninfo string

func InitMySql(sqlurl, sqluser, sqlpass, sqldb string) (db *gorm.DB, err error) {

	// 连接数据库
	conninfo = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sqluser, sqlpass, sqlurl, sqldb)
	log.Println(conninfo)

	db, err = gorm.Open("mysql", conninfo)
	//defer db.Close()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// 创建必须的数据库表
	// 1. users 表
	if db.HasTable(User{}) == false {
		err = db.CreateTable(User{}).Error
		if err != nil {
			log.Println(err)
			panic(err)
		} else {
			log.Println("Create users table success")
		}
	} else {
		log.Println("users table exist!")
	}

	// 2. friends 表
	if db.HasTable(Friend{}) == false {
		err = db.CreateTable(Friend{}).Error
		if err != nil {
			log.Println(err)
			panic(err)
		} else {
			log.Println("Create friends table success")
		}
	} else {
		log.Println("friends table exist!")
	}

	// 3. diss 表
	if db.HasTable(Diss{}) == false {
		err = db.CreateTable(Diss{}).Error
		if err != nil {
			log.Println(err)
			panic(err)
		} else {
			log.Println("Create diss table success")
		}
	} else {
		log.Println("diss table exist!")
	}


	return db, nil
}

// models 包初始化，创建2个 db 对象并且读取配置
// db1 为 gorm 对象，适合构建表
// db2 为 sql 驱动对象，轻量一些 一些小 sql 语句直接执行即可
var db *gorm.DB
var db2 *sql.DB

func init()  {
	// 进入 models 包先执行 init() 读数据库配置并连接初始化

	// 从配置文件 conf/app.conf 中读取配置
	sqluser := beego.AppConfig.String("mysqluser")
	sqlpass := beego.AppConfig.String("mysqlpass")
	sqlurl  := beego.AppConfig.String("mysqlurls")
	sqldb   := beego.AppConfig.String("mysqldb")

	var err error
	db, err = InitMySql(sqlurl, sqluser, sqlpass, sqldb)
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("db init success")
	}

	db2, err = sql.Open("mysql", conninfo)
	if err != nil {
		log.Println("db2 connect fail")
		log.Println(err)
	} else {
		log.Println("db2 init success")
	}
}

