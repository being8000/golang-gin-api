package service_test

import (
	"fmt"
	"log"
	"testing"
	"zehan/gin/model"
	"zehan/gin/service"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Role = model.Role

func connectDB() {
	config := map[string]string{
		"user":     "root",
		"password": "123456",
		"host":     "localhost",
		"port":     "3306",
		"name":     "test",
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["port"],
		config["name"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// PrepareStmt: true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println(err)
	}
	DB = db

	DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Menu{},
	)
}

type Menu = model.Menu

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}
func TestAddMenus(t *testing.T) {
	connectDB()
	menu := &Menu{Name: "User Management", NodeType: 1, Code: "user"}
	mres := DB.Create(menu)
	if mres.Error != nil {
		log.Panic(mres.Error)
	}

	menus := []*Menu{
		{Name: "Add", NodeType: 2, Code: "add", ParentID: menu.ID},
		{Name: "Edit", NodeType: 2, Code: "edit", ParentID: menu.ID},
		{Name: "Delete", NodeType: 2, Code: "delete", ParentID: menu.ID},
		{Name: "List", NodeType: 2, Code: "list", ParentID: menu.ID},
		{Name: "Export", NodeType: 2, Code: "export", ParentID: menu.ID},
		{Name: "Export2", NodeType: 2, Code: "export2", ParentID: menu.ID},
	}
	mres = DB.Create(menus)
	if mres.Error != nil {
		t.Error(mres.Error)
	}
	menus = append(menus, menu)

	role := []*Role{
		{Name: "Admin", Menus: menus},
		{Name: "Role1"},
		{Name: "Role2"},
	}

	result := DB.Create(role)
	if result.Error != nil {
		t.Error(result.Error)
	} else {
		t.Log(result)
	}
}

func TestBindRoleAndMenu(t *testing.T) {
	connectDB()
	menu := &Menu{Name: "User Management", NodeType: 1, Code: "user"}
	result := DB.Create(menu)
	if result.Error != nil {
		t.Error(result.Error)
	}

	menus := []*Menu{
		{Name: "Add", NodeType: 2, Code: "add", ParentID: menu.ID},
		{Name: "Edit", NodeType: 2, Code: "edit", ParentID: menu.ID},
		{Name: "Delete", NodeType: 2, Code: "delete", ParentID: menu.ID},
		{Name: "List", NodeType: 2, Code: "list", ParentID: menu.ID},
		{Name: "Export", NodeType: 2, Code: "export", ParentID: menu.ID},
	}
	result = DB.Create(menus)
	if result.Error != nil {
		t.Error(result.Error)
	}
}

type User = model.User

func TestBindRole2User(t *testing.T) {
	connectDB()
	user := &User{}
	result := DB.Model(&User{ID: 1}).First(user)
	if result.Error != nil {
		t.Error(result.Error)
	}
	roles := []*Role{}
	result = DB.Where("name = ?", "Admin").Find(&roles)
	log.Print("Find", result.RowsAffected)
	if result.Error != nil {
		t.Error(result.Error)
	}
	// user.Roles = roles
	aso := DB.Model(&user).Association("Roles")
	if aso.Error != nil {
		t.Error(aso.Error)
	}

	aso.Append(roles)
	// log.Print("Save", result.RowsAffected)
	// if result.Error != nil {
	// 	t.Error(result.Error)
	// }
}

func TestAdd(t *testing.T) {
	if ans := service.Add(1, 4); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := service.Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}
