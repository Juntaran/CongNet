/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/20 12:16 
  */

package models

import (
	"time"
	"log"
	"errors"
)

type Diss struct {
	ID          	uint    	`gorm:"primary_key;AUTO_INCREMENT"`		// ID 主键
	AutherID		uint		`gorm:"not null"`						// 作者ID
	AuthorName		string  	`gorm:"type:varchar(128);not null"` 	// 作者名
	Content			string		`gorm:"type:varchar(256);not null"`		// 内容			长度限制为 256 一个汉字占3个字符，英文占1个	select * from disses WHERE LENGTH(content)=24
	CreateTime 		time.Time	`gorm:"not null"` 						// 发布时间
	Report 			string		`gorm:"type:varchar(256)"`				// 转发附议
	Original		string		`gorm:"type:varchar(128)"`				// 初始 dissID
}

// 发布一个 Diss
func CreateDiss(diss Diss) error {
	err := db.Create(&diss).Error
	return err
}

// 删除一个 Diss
func DeleteDiss(dissID string) error {

	// 判断这个 dissID 是否存在
	var content string
	err := db2.QueryRow("SELECT content FROM disses WHERE id=?", dissID).Scan(&content)
	if err != nil || content == "" {
		return errors.New(dissID + "not Exist")
	}

	result, err := db2.Exec("DELETE FROM disses WHERE id=?", dissID)
	if err != nil {
		log.Println("Delete diss Fail")
		log.Println(err)
		return err
	} else {
		log.Println("Delete diss Success")
		log.Println(result)
		return nil
	}
}

// 根据 dissID 查找发布用户的 ID
func GetUserIDByDissID(dissID string) string {
	var userID string
	db2.QueryRow("SELECT auther_id FROM disses WHERE id=?", dissID).Scan(&userID)
	log.Println("Get", userID, "By", dissID)
	return userID
}

// 根据 dissID 查找发布用户的 email
func GetEmailByDissID(dissID string) string {
	userID := GetUserIDByDissID(dissID)
	var email string
	db2.QueryRow("SELECT email FROM users WHERE id=?", userID).Scan(&email)
	log.Println("Get", email, "By", dissID)
	return email
}

// 根据 dissID 查找 diss 内容
func GetDissContentByDissID(dissID string) string {
	var content string
	db2.QueryRow("SELECT content FROM disses WHERE id=?", dissID).Scan(&content)
	log.Println("Get", content, "By", dissID)
	return content
}

// 转发一个 diss
func ReportDissByDissID(diss Diss) error {
	// 重新创建一个 	diss
	// AutherID 	替换成转发用户的 ID
	// AutherName	替换成转发用户的 Name
	// Content		不变
	// CreateTime	替换为当前时间
	// Report		增加转发评语
	err := db.Create(&diss).Error
	return err
}

// 根据 dissID 查找 dissRep
func GetDissRepByDissID(dissID string) string {
	var content string
	db2.QueryRow("SELECT report FROM disses WHERE id=?", dissID).Scan(&content)
	log.Println("Get", content, "By", dissID)
	return content
}

// 根据 dissID 查找 original
func GetDissOriByDissID(dissID string) string {
	var original string
	db2.QueryRow("SELECT original FROM disses WHERE id=?", dissID).Scan(&original)
	log.Println("Get", original, "By", dissID)
	return original
}