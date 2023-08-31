package pkg

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:50;uniqueIndex:unique_index"`
	V0    string `gorm:"size:50;uniqueIndex:unique_index"`
	V1    string `gorm:"size:50;uniqueIndex:unique_index"`
	V2    string `gorm:"size:50;uniqueIndex:unique_index"`
	V3    string `gorm:"size:50;uniqueIndex:unique_index"`
	V4    string `gorm:"size:50;uniqueIndex:unique_index"`
	V5    string `gorm:"size:50;uniqueIndex:unique_index"`
}

type Casbin struct {
	Adapter *gormadapter.Adapter
	R1      *casbin.Enforcer // RBAC MODEL
}

func NewCasbin(db *gorm.DB) *Casbin {
	adapter := getCasbinAdapter(db)
	return &Casbin{
		Adapter: adapter,
	}
}

func getCasbinAdapter(db *gorm.DB) *gormadapter.Adapter {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	if err != nil {
		panic(err)
	}
	return a
}

func getCasbinRBACEnforcer(a *gormadapter.Adapter) *casbin.Enforcer {
	e, err := casbin.NewEnforcer("./app/conf/rbac_model.conf", a)
	if err != nil {
		panic(err)
	}
	// Load the policy from DB.
	e.LoadPolicy()
	return e
}

func GetEnforcer(policy string) *casbin.Enforcer {
	e, err := casbin.NewEnforcer("./app/conf/rbac_model.conf", "./app/casbin_policy/"+policy)
	if err != nil {
		log.Panicf("NewEnforecer failed:%v\n", err)
	}
	return e
}
