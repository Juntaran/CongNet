/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/14 18:03 
  */

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

// mysql 数据库初始化
func InitMySql(sqlurl, sqluser, sqlpass, sqldb string) (db *gorm.DB, err error) {
	var conninfo string = sqluser + ":" + sqlpass + "@tcp(" + sqlurl + ")/" + sqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(conninfo)
	db, err = gorm.Open("mysql", conninfo)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return db, err
	}
	return db, nil
}
