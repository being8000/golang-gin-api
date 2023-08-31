package service

import (
	"fmt"
	"zehan/gin/handlers/vo"
	"zehan/gin/model"
	"zehan/gin/utils"
)

type User = model.User
type UserService struct {
	Kit *utils.Kit
}

func (s UserService) Login(params *vo.LoginForm) (user *User, err error) {
	db := s.Kit.DB
	r := db.Where("mobile = ?", params.Mobile).First(user)
	if r.RowsAffected == 0 {
		err = fmt.Errorf("Account not found, please go to register first.")
		return
	}
	err = utils.ComparePassword(user.Password, params.Password)
	if err != nil {
		return
	}

	return
}

func (s UserService) Add(vUser *vo.User) error {
	// user, err := json.Marshal(vUser)
	// // address, err := json.Marshal(vUser.Addresses)
	// if err != nil {
	// 	return err
	// }
	// json.Unmarshal([]byte(user), &s.userDTO)
	// // json.Unmarshal([]byte(address), &s.userDTO.Addresses)
	// return s.DB.Create(s.userDTO).Error
	return nil
}

func (s UserService) Get(id string) (user model.User, err error) {
	s.Kit.DB.Preload("Addresses").First(&user, id)
	return
}
