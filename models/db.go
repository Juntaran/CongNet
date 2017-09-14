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
)

// mysql 数据库初始化
// 1. 连接数据库
// 2. 创建必须的数据库表
func InitMySql(sqlurl, sqluser, sqlpass, sqldb string) (db *gorm.DB, err error) {

	// 连接数据库
	conninfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sqluser, sqlpass, sqlurl, sqldb)
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
			log.Println("Create table success")
		}
	} else {
		log.Println("users table exist!")
	}



	return db, nil
}
