/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/21 10:14 
  */

package models

import (
	"time"
	"errors"
	"log"
)

type Comment struct {
	ID          	uint    	`gorm:"primary_key;AUTO_INCREMENT"`		// ID 主键
	AutherID		uint		`gorm:"not null"`						// 作者ID
	AuthorName		string  	`gorm:"type:varchar(128);not null"` 	// 作者名
	CommentType		uint		`gorm:"not null"`						// 评论的类型，是 diss 还是日志 还是 相册，目前 diss 对应0
	TypeID			string		`gorm:"not null"`						// 对应的ID 比如是对 dissID为100的 diss 的评论，这个字段记录100
	CreateTime 		time.Time	`gorm:"not null"` 						// 发布时间
	CommentContent	string		`gorm:"type:varchar(256);not null"`		// 内容	长度限制为 256 一个汉字占3个字符，英文占1个
}

// 发布一个 Comment
func CreateComment(comment Comment) error {
	err := db.Create(&comment).Error
	return err
}

// 删除一个 Comment
func DeleteComment(commentID string) error {

	// 判断这个 commentID 是否存在
	var content string
	err := db2.QueryRow("SELECT comment_content FROM comments WHERE id=?", commentID).Scan(&content)
	if err != nil || content == "" {
		return errors.New(commentID + "not Exist")
	}

	result, err := db2.Exec("DELETE FROM comments WHERE id=?", commentID)
	if err != nil {
		log.Println("Delete comment Fail")
		log.Println(err)
		return err
	} else {
		log.Println("Delete comment Success")
		log.Println(result)
		return nil
	}
}

// 根据 commentID 查找发布用户的 ID
func GetUserIDByCommentID(commentID string) string {
	var userID string
	db2.QueryRow("SELECT auther_id FROM comments WHERE id=?", commentID).Scan(&userID)
	log.Println("Get", userID, "By", commentID)
	return userID
}

// 根据 commentID 查找发布用户的 email
func GetEmailByCommentID(commentID string) string {
	userID := GetUserIDByCommentID(commentID)
	var email string
	db2.QueryRow("SELECT email FROM users WHERE id=?", userID).Scan(&email)
	log.Println("Get", email, "By", commentID)
	return email
}

// 根据 dissID 查找所有的 comment
func GetCommentByDissID(dissID string) []string {
	// 结果集合
	var comments []string

	// select * from comments where u_id = userid or f_id = usedrid
	rows, err := db2.Query("SELECT comment_content FROM comments WHERE type_id=? AND comment_type=?", dissID, "0")
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var retFid string
			if err = rows.Scan(&retFid); err == nil {
				log.Println(err)
			}
			comments = append(comments, retFid)
		}
	}
	return comments
}