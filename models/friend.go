/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/18 10:43 
  */

package models

import (
	"log"
	"errors"
)

type Friend struct {
	//gorm.Model	// gorm 会自动设置字段 ID 默认为主键自增，也会设置软删除相关字段
	ID          	uint    `gorm:"primary_key;AUTO_INCREMENT"`
	U_ID			uint												// 发起好友请求的用户，UserID 外键
	F_ID			uint												// 接受好友请求的用户，UserID 外键
	Content 		string	`gorm:"type:varchar(128)"`					// 申请留言
	ACCEPT			uint												// 是否接受该请求，0为未接受，1为接受，2为拒绝，3为拉黑
}



// 判断两个人是否为好友
func IsFriend( id1, id2 uint) bool {

	var judge bool
	if db.Where("U_ID = ? AND F_ID = ? AND ACCEPT = ?", id1, id2, 1) == nil {
		judge = false
	} else {
		return true
	}
	if db.Where("U_ID = ? AND F_ID = ? AND ACCEPT = ?", id2, id1, 1) == nil {
		judge = false
	} else {
		return true
	}
	return judge
}

// 发起一个添加好友请求
func StartAddFriend(u_id, f_id uint, content string) error {
	friend := &Friend{
		U_ID:   		u_id,
		F_ID: 			f_id,
		Content: 		content,
		ACCEPT: 		0,
	}
	err := db.Create(&friend).Error
	return err
}

// 通过一个好友请求
func AcceptAddFriend(u_id, f_id uint) error {
	friend := &Friend{}
	ret := db.Where("U_ID = ? AND F_ID = ? AND ACCEPT = ?", u_id, f_id, 0).First(&friend)
	err := ret.Error
	if err == nil {
		friend.ACCEPT = 1
		err2 := ret.Select("accept").Updates(friend).Error
		if err2 != nil {
			log.Println(err2)
			log.Println(u_id, f_id, "Accept Add Friend Error")
			return err2
		} else {
			log.Println(u_id, f_id, "Accept Add Friend Success")
			return nil
		}
	} else {
		log.Println(err)
		log.Println("Can't find this Request")
		return err
	}
}

// 拒绝一个好友请求
func RefuseAddFriend(u_id, f_id uint) error {
	friend := &Friend{}
	ret := db.Where("U_ID = ? AND F_ID = ? AND ACCEPT = ?", u_id, f_id, 0).First(&friend)
	err := ret.Error
	if err == nil {
		friend.ACCEPT = 2
		err2 := ret.Select("accept").Updates(friend).Error
		if err2 != nil {
			log.Println(err2)
			log.Println(u_id, f_id, "Refuse Add Friend Error")
			return err2
		} else {
			log.Println(u_id, f_id, "Refuse Add Friend Success")
			return nil
		}
	} else {
		log.Println(err)
		log.Println("Can't find this Request")
		return err
	}
}

// 拉黑一个好友
func BlockAddFriend(u_id, f_id uint) error {
	friend := &Friend{}
	ret := db.Where("U_ID = ? AND F_ID = ?", u_id, f_id).First(&friend)
	err := ret.Error
	if err == nil {
		friend.ACCEPT = 3
		err2 := ret.Select("accept").Updates(friend).Error
		if err2 != nil {
			log.Println(err2)
			log.Println(u_id, f_id, "Block Friend Error")
			return err2
		} else {
			log.Println(u_id, f_id, "Block Friend Success")
			return nil
		}
	} else {
		log.Println(err)
		log.Println("Can't find this Request")
		return err
	}
}

// 获取一个用户的所有好友，返回他的好友id list 以及 list 大小
func GetAllFriend(userid uint) ([]uint, int, error) {
	// 先判断这个 userid 是否存在
	//ret := db.Where("id = ?", userid).First(&User{}).Scan(&User{})
	err := db2.QueryRow("SELECT * FROM users WHERE id=?", userid).Scan(&User{})
	if err != nil {
		log.Println("This User not Exist")
		return nil, 0, errors.New("This User not Exist")
	}
	if err == nil {
		// 是否为注销用户判断
		isCancel := &User{
			ID: 	userid,
		}
		db.Where("id = ?", userid).First(&User{}).Scan(&User{})
		db.First(&isCancel,"id = ?", userid)
		if isCancel.Cancel == 1 {
			log.Println("Farewell, this user is Canceled")
			err = errors.New("Farewell, this user is Canceled")
			return nil, 0, err
		}
		log.Println("Start Get", userid, "Friends")

		// 结果集合
		var friends []uint

		// select * from friends where u_id = userid or f_id = usedrid
		rows1, err := db2.Query("SELECT f_id FROM friends WHERE u_id=?", userid)
		if err != nil {
			log.Println(err)
		} else {
			for rows1.Next() {
				var retFid uint
				if err = rows1.Scan(&retFid); err == nil {
					log.Println(err)
				}
				friends = append(friends, retFid)
			}
		}

		rows2, err := db2.Query("SELECT u_id FROM friends WHERE f_id=?", userid)
		if err != nil {
			log.Println(err)
		} else {
			for rows2.Next() {
				var retUid uint
				if err = rows2.Scan(&retUid); err == nil {
					log.Println(err)
				}
				friends = append(friends, retUid)
			}
		}

		log.Println("Search Success")
		return friends, len(friends), errors.New("Search Success")
	} else {
		log.Println("Search Fail")
		return nil, 0, errors.New("Search Fail")
	}
}