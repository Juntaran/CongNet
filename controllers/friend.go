/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/18 15:46 
  */

package controllers

import (
	"CongNet/models"
	"strconv"
)

// 查询好友
type GetFriendsController struct {
	BaseController
}

func (this *GetFriendsController) Get() {
	check := this.isLogin
	if !check {
		this.Redirect("/login", 302)
	}
}

func (this *GetFriendsController) Post() {

	userget := this.GetString("userid")
	usergetInt, err := strconv.Atoi(userget)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败"}
		this.ServeJSON()
		return
	}
	userid := uint(usergetInt)

	friends, length, err := models.GetAllFriend(userid)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "您有"+strconv.Itoa(length)+"个好友": friends}
	this.ServeJSON()
}
