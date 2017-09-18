/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/14 18:03 
  */

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
	"github.com/astaxie/beego"
	"database/sql"
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

	return db, nil
}

// models 包初始化，创建一个 db 对象并且读取配置
var db *gorm.DB
var db2 *sql.DB

func init()  {
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

