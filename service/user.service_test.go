package service_test

import (
	"encoding/json"
	"log"
	"testing"
	"zehan/gin/handlers/vo"
	"zehan/gin/model"
	"zehan/gin/utils"
)

type VoUser = vo.User

func TestUserAdd(t *testing.T) {
	vUser := &VoUser{
		FirstName: "elon",
		LastName:  "chen",
		Age:       19,
		Password:  "7086052czh",
		Mobile:    15980271371,
	}
	cryptPassword := utils.HashPassword(vUser.Password)
	connectDB()
	result := DB.Create(&model.User{
		FirstName: vUser.FirstName,
		LastName:  vUser.LastName,
		Age:       vUser.Age,
		Password:  cryptPassword,
		Mobile:    vUser.Mobile,
	})
	log.Println(result)
}

func TestComparePassword(t *testing.T) {
	connectDB()
	user := &model.User{}
	r := DB.Where("mobile = ?", 15980271371).First(user)
	log.Println(r)
	b := utils.ComparePassword(user.Password, "7086052czh123")
	log.Println(b)
}

func TestGetUser(t *testing.T) {
	connectDB()
	user := &model.User{}
	DB.Debug().Preload("Roles").Preload("Roles.Menus").Preload("Children").First(user, "m_users.mobile = ?", 15980271371)
	// ary := make([]uint, 0)
	// for _, v := range user.Roles {
	// 	ary = append(ary, v.ID)
	// }
	// log.Println(ary)
	// DB.Debug().Preload("Menus", "role_id in (?)", ary).Find(&user.Roles)
	byt, _ := json.Marshal(user)
	log.Println(string(byt))
}
