/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 0:13
 */

package models

import (
	"log"
	"errors"
)

type User struct {
	//gorm.Model	// gorm 会自动设置字段 ID 默认为主键自增，也会设置软删除相关字段
	ID          	uint    `gorm:"primary_key;AUTO_INCREMENT"` 				// ID 主键
	Name			string  `gorm:"type:varchar(128);unique_index;not null"` 	// string 默认长度为255, 使用这种tag重设。
	Password		string	`gorm:"type:varchar(128);not null"`					// 密码
	Email 			string	`gorm:"type:varchar(128);unique_index;not null"` 	// `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Cancel 			uint														// 是否注销 0为正常用户，1为已经注销用户
}

// 用户注册，传入 db 指针 和 用户信息，返回 error
func RegisterUser(user User) error {
	err := db.Create(&user).Error
	return err
}

// 根据用户邮箱查找用户ID
func GetIDByEmail(email string) uint {
	// 这个 email 是一定存在的，所以懒得返回 error 了
	var id uint
	db2.QueryRow("SELECT id FROM users WHERE email=?", email).Scan(&id)
	log.Println("Get", id, "By", email)
	return id
}

// 用户登陆，传入 db 指针 和 用户信息，返回 error
func LoginUser(email, password string) (int, error) {

	log.Println("start login", email, password)
	ret := db.Where("email = ? AND password = ?", email, password).First(&User{}).Scan(&User{})
	log.Println(db.Where("email = ? AND password = ?", email, password).First(&User{}))

	err := ret.Error
	if err == nil {

		// 是否为注销用户判断
		isCancel := &User{
			Email: 		email,
			Password: 	password,
		}
		db.First(&isCancel,"email = ? AND password = ?", email, password)
		if isCancel.Cancel == 1 {
			log.Println("Farewell, this user is Canceled")
			err = errors.New("Farewell, this user is Canceled")
			return 1, err
		}
		log.Println(email, "Login Success")
	} else {
		log.Println(email, "Login Fail")
	}
	return 0, err
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

// 删除用户，不会真正把用户删掉，会把密码清空
func DeleteUser(name, password, email string) error {
	user := &User{}
	ret := db.Where("name = ? AND password = ? AND email = ?", name, password, email).First(&User{})
	err := ret.Error
	if err == nil {
		ret.First(&user, "name=?", name)
		user.Cancel = 1
		err2 := ret.Select("cancel").Updates(user).Error
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