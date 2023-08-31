package service

import (
	"fmt"
	"zehan/gin/model"

	"gorm.io/gorm"
)

type Role = model.Role
type PermissionService struct {
	DB *gorm.DB
}

func (s PermissionService) AddRole() {
	role := []*Role{
		{Name: "Test"},
		{Name: "Test2"},
	}
	result := s.DB.Create(role)
	fmt.Println(result.RowsAffected)
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}
