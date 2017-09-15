/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 0:13
 */

package models

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type User struct {
	//gorm.Model
	ID          	uint    `gorm:"primary_key;AUTO_INCREMENT"` 				// gorm 会自动设置字段 ID 默认为主键自增
	Name			string  `gorm:"type:varchar(128);unique_index;not null"` 	// string 默认长度为255, 使用这种tag重设。
	Password		string	`gorm:"type:varchar(128);not null"`					// 密码
	Email 			string	`gorm:"type:varchar(128);unique_index;not null"` 	// `type`设置sql类型, `unique_index` 为该列设置唯一索引
}

var db *gorm.DB

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
}

// 用户注册，传入 db 指针 和 用户信息，返回 error
func RegisterUser(user User) error {
	err := db.Create(&user).Error
	return err
}

// 用户登陆，传入 db 指针 和 用户信息，返回 error
func LoginUser(name, password string) error {

	log.Println("start login", name, password)
	ret := db.Where("name = ? AND password = ?", name, password).First(&User{}).Scan(&User{})
	log.Println(db.Where("name = ? AND password = ?", name, password).First(&User{}))
	err := ret.Error
	if err == nil {
		log.Println(name, "Login Success")
	} else {
		log.Println(name, "Login Fail")
	}
	return err
}

// 根据 Email 查找用户信息
func SearchUserByEmail(email string) (string, error) {
	user := &User{
		Email: 	email,
	}
	ret := db.First(&user, "email=?", email)
	err := ret.Error
	var rets string
	if err != nil {
		rets = "Your Email Wrong"
		return rets, err
	}
	rets = "\nUsername: " + user.Name + "\nPassword: " + user.Password + "\nEmail: " + user.Email

	return rets, err
}

// 删除用户  软删除，在数据库中记录删除时间，不会真正删除记录
func DeleteUser(name, password string) error {
	ret := db.Delete(User{}, "name = ? AND password = ?", name, password)
	err := ret.Error
	if err == nil {
		log.Println("Delete Success")
	} else {
		log.Println("Delete Fail")
	}
	return err
}

// 修改密码
func UpdateUserPassword(name, password, email, newpassword string) error {
	user := &User{}
	ret := db.Where("name = ? AND password = ? AND email = ?", name, password, email).First(&User{})
	err := ret.Error
	if err == nil {
		ret.First(&user, "name=?", name)
		user.Password = newpassword
		err2 := ret.Select("password").Updates(user).Error
		if err2 != nil {
			log.Println(err2)
			log.Println("Update Password Error")
			return err2
		} else {
			log.Println("Update Password Success")
			return nil
		}
	} else {
		log.Println(err)
		log.Println("Your Information Wrong")
		return err
	}
}

// 修改邮箱
func UpdateUserEmail(name, password, newemail string) error {
	user := &User{}
	ret := db.Where("name = ? AND password = ?", name, password).First(&User{})
	err := ret.Error
	if err == nil {
		ret.First(&user, "name=?", name)
		user.Email = newemail
		err2 := ret.Select("email").Updates(user).Error
		if err2 != nil {
			log.Println(err2)
			log.Println("Update Email Error")
			return err2
		} else {
			log.Println("Update Email Success")
			return nil
		}
	} else {
		log.Println(err)
		log.Println("Your Information Wrong")
		return err
	}
}