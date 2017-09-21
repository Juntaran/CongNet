/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/20 16:00 
  */

package controllers

import (
	"CongNet/models"
	"time"
	"log"
)

// create comment
type CreateCommentController struct {
	BaseController
}

func (this *CreateCommentController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "comment.tpl"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *CreateCommentController) Post() {
	// id, author_id, author_name, comment_type, create_time, comment_content
	log.Println("Into Create Comment Controller")
	author_email := this.userEmail
	author_id := models.GetIDByEmail(author_email)
	author_name := models.GetNameByEmail(author_email)
	var comment_type uint = 0
	type_id := this.GetString("typeid")
	create_time := time.Now()
	comment_content := this.GetString("content")

	log.Println("author_email:", author_email)
	log.Println("author_name:", author_name)

	if comment_content == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "吐槽不能为空哟~"}
		this.ServeJSON()
	}

	comment := models.Comment{
		AutherID: 			author_id,
		AuthorName: 		author_name,
		CommentType:		comment_type,
		TypeID: 			type_id,
		CreateTime: 		create_time,
		CommentContent: 	comment_content,
	}
	log.Println(comment)
	err := models.CreateComment(comment)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "评论发布成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "评论发布失败"}
	}
	this.ServeJSON()
}


// delete comment
type DeleteCommentController struct {
	BaseController
}

func (this *DeleteCommentController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "commentDel.tpl"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *DeleteCommentController) Post() {
	commentID := this.GetString("commentID")

	// 判断这个 comment 是不是当前用户发布的
	authorEmail := this.userEmail		// 当前用户ID
	userEmail := models.GetEmailByCommentID(commentID)
	log.Println("authorEmail", authorEmail)
	log.Println("userEmail", userEmail)
	if userEmail != authorEmail {
		log.Println("userEmail != authorEmail")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "你没有权限哦~"}
		this.ServeJSON()
		return
	}

	if commentID == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "commentID is nil~"}
		this.ServeJSON()
		return
	}

	err := models.DeleteComment(commentID)
	if err != nil {
		log.Println("Delete", commentID, "Fail")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": commentID + " not Exist"}
		this.ServeJSON()
	} else {
		log.Println("Delete", commentID, "Success")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Delete " + commentID + " Success"}
		this.ServeJSON()
	}
}