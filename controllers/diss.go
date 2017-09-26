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

// create diss
type CreateDissController struct {
	BaseController
}

func (this *CreateDissController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "diss.html"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *CreateDissController) Post() {
	// id, author_id, author_name, content, create_time, original
	author_email := this.userEmail
	author_id := models.GetIDByEmail(author_email)
	author_name := models.GetNameByEmail(author_email)
	content := this.GetString("content")
	create_time := time.Now()

	if content == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "吐槽不能为空哟~"}
		this.ServeJSON()
	}

	diss := models.Diss{
		AutherID: 			author_id,
		AuthorName: 		author_name,
		Content: 			content,
		CreateTime: 		create_time,
		Original: 			"0",
	}
	err := models.CreateDiss(diss)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "吐槽发布成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "吐槽发布失败"}
	}
	this.ServeJSON()
	// 跳转到 / ，如果用户已经登录，会自动跳转到 /user/userid ，相当于刷新当前页
	this.Redirect("/", 302)
}


// delete diss
type DeleteDissController struct {
	BaseController
}

func (this *DeleteDissController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "dissDel.html"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *DeleteDissController) Post() {
	dissID := this.GetString("dissID")

	// 判断这个 diss 是不是当前用户发布的
	authorEmail := this.userEmail		// 当前用户ID
	userEmail := models.GetEmailByDissID(dissID)
	log.Println("authorEmail", authorEmail)
	log.Println("userEmail", userEmail)
	if userEmail != authorEmail {
		log.Println("userEmail != authorEmail")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "你没有权限哦~"}
		this.ServeJSON()
		return
	}

	if dissID == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "dissID is nil~"}
		this.ServeJSON()
		return
	}

	err := models.DeleteDiss(dissID)
	if err != nil {
		log.Println("Delete", dissID, "Fail")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": dissID + " not Exist"}
		this.ServeJSON()
	} else {
		log.Println("Delete", dissID, "Success")
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Delete " + dissID + " Success"}
		this.ServeJSON()
	}
}


// report diss
type ReportDissController struct {
	BaseController
}

func (this *ReportDissController) Get() {
	check := this.isLogin
	if check {
		this.TplName = "dissRep.html"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *ReportDissController) Post() {
	// id, author_id, author_name, content, create_time
	// 重新创建一个 	diss
	// AutherID 	替换成转发用户的 ID
	// AutherName	替换成转发用户的 Name
	// Content		不变
	// CreateTime	替换为当前时间
	// Report		增加转发评语

	dissID := this.GetString("dissID")
	oldReport := models.GetDissRepByDissID(dissID)

	author_email := this.userEmail
	author_id := models.GetIDByEmail(author_email)
	author_name := models.GetNameByEmail(author_email)
	content := models.GetDissContentByDissID(dissID)
	create_time := time.Now()
	var original string
	var judge int = 0

	if len(oldReport) == 0 {
		// 第一个转发
		original = dissID
		judge = 1
	} else {
		// 非第一个转发
		// original 保持不变
		original = models.GetDissOriByDissID(dissID)
	}
	newReport := this.GetString("report")
	var report string = author_name + "：" + newReport + "//" + oldReport

	if len(newReport) == 0 {
		// 不添加评论
		// log.Println("do nothing")
	} else {
		// 如果是不是第一个转发，则给最初的作者评论
		if judge == 0 {
			// log.Println("给最初的作者评论")
			comment1 := models.Comment{
				AutherID: 			author_id,
				AuthorName: 		author_name,
				CommentType:		0,
				TypeID: 			original,
				CreateTime: 		create_time,
				CommentContent: 	newReport,
			}
			log.Println(comment1)
			models.CreateComment(comment1)
		}
		log.Println("给上一个转发者评论")
		// 把 newReport 写入转发的 comment
		comment2 := models.Comment{
			AutherID: 			author_id,
			AuthorName: 		author_name,
			CommentType:		0,
			TypeID: 			dissID,
			CreateTime: 		create_time,
			CommentContent: 	newReport,
		}
		log.Println(comment2)
		models.CreateComment(comment2)
	}

	// 转发
	diss := models.Diss{
		AutherID: 			author_id,
		AuthorName: 		author_name,
		Content: 			content,
		CreateTime: 		create_time,
		Report:				report,
		Original: 			original,
	}
	err := models.CreateDiss(diss)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "吐槽转发成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "吐槽转发失败"}
	}
	this.ServeJSON()
	// 跳转到 / ，如果用户已经登录，会自动跳转到 /user/userid ，相当于刷新当前页
	//this.Redirect("/", 302)
}