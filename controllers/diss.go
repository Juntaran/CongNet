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
		this.TplName = "diss.tpl"
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *CreateDissController) Post() {
	// id, author_id, author_name, content, create_time
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
		this.TplName = "dissDel.tpl"
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