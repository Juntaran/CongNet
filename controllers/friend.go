/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/18 15:46 
  */

package controllers

import (
	"CongNet/models"
	"strconv"
	"log"
)

// 查询好友
type FriendsGetController struct {
	BaseController
}

func (this *FriendsGetController) Get() {
	check := this.isLogin
	if !check {
		this.Redirect("/login", 302)
	} else {
		this.TplName = "friends.html"
	}
}

func (this *FriendsGetController) Post() {

	userid := this.GetString("userid")
	friends, length, err := models.GetAllFriend(userid)
	if err == nil {
		// 格式化返回字符串
		var ret string = "您有" + strconv.Itoa(length) + "个好友：\n"
		for i := 0; i < length - 1; i++ {
			ret += friends[i]
			ret += "\n"
		}
		ret += friends[length-1]
		log.Println(ret)

		this.Data["json"] = map[string]interface{}{"code": 1, "message": ret}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败"}
		this.ServeJSON()
	}
}
