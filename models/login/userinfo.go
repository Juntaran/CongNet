/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/14 18:39 
  */

package login

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID          	int     `gorm:"AUTO_INCREMENT"` 				// 自增，字段 ID 默认为主键
	Name			string  `gorm:"type:varchar(100);unique_index"` // string 默认长度为255, 使用这种tag重设。
	Password		string	`gorm:"type:varchar(100)"`				// 密码
	Email 			string	`gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
}