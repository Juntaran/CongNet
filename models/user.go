/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/15 0:13
 */

package models

//import "github.com/jinzhu/gorm"

type User struct {
	//gorm.Model
	ID          	uint    `gorm:"primary_key;AUTO_INCREMENT"` 				// gorm 会自动设置字段 ID 默认为主键自增
	Name			string  `gorm:"type:varchar(128);unique_index;not null"` 	// string 默认长度为255, 使用这种tag重设。
	Password		string	`gorm:"type:varchar(128);not null"`					// 密码
	Email 			string	`gorm:"type:varchar(128);unique_index;not null"` 	// `type`设置sql类型, `unique_index` 为该列设置唯一索引
}