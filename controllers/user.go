/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 1:25
 */

package controllers

import (
	"github.com/jinzhu/gorm"
	"CongNet/models"
	"log"
)

// 用户注册，传入 db 指针 和 用户信息，返回 error
func RegisterUser(db *gorm.DB, user models.User) error {
	err := db.Create(&user).Error
	return err
}

// 用户登陆，传入 db 指针 和 用户信息，返回 error
func LoginUser(db *gorm.DB, name, password string) error {
	ret := db.Where("name = ? AND password = ?", name, password).First(&models.User{}).Scan(&models.User{})
	err := ret.Error
	if err == nil {
		log.Println(name, "Login Success")
	} else {
		log.Println(name, "Login Fail")
	}
	return err
}

// 根据 Email 查找用户信息
func SearchUserByEmail(db *gorm.DB, email string) (string, error) {
	user := &models.User{
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
func DeleteUser(db *gorm.DB, name, password string) error {
	ret := db.Delete(models.User{}, "name = ? AND password = ?", name, password)
	err := ret.Error
	if err == nil {
		log.Println("Delete Success")
	} else {
		log.Println("Delete Fail")
	}
	return err
}

// 修改密码
func UpdateUserPassword(db *gorm.DB, name, password, email, newpassword string) error {
	user := &models.User{}
	ret := db.Where("name = ? AND password = ? AND email = ?", name, password, email).First(&models.User{})
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
func UpdateUserEmail(db *gorm.DB, name, password, newemail string) error {
	user := &models.User{}
	ret := db.Where("name = ? AND password = ?", name, password).First(&models.User{})
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