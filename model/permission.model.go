package model

type Role struct {
	ID        uint     `json:"id"`
	Name      string   `gorm:"uniqueIndex;type:varchar(50);not null;"`
	Menus     []*Menu  `gorm:"many2many:p_role_memus;"`
	Users     []*User  `json:"-" gorm:"many2many:p_user_roles;"`
	CreatedAt NullTime `json:"createdAt"`
	UpdatedAt NullTime `json:"updatedAt"`
	DeletedAt NullTime `json:"deletedAt" gorm:"index"`
}

type Menu struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name" gorm:"type:varchar(50);uniqueIndex:idx_names"`
	NodeType  uint8    `json:"nodeType" gorm:"type:tinyint;comment:1=Page, 2=Button"`
	Code      string   `json:"code" gorm:"type:varchar(50);"`
	ParentID  uint     `json:"parentId" gorm:"uniqueIndex:idx_names"`
	Parent    *Menu    `json:"-"`
	Children  []*Menu  `json:"children,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	CreatedAt NullTime `json:"-"`
	UpdatedAt NullTime `json:"-"`
	DeletedAt NullTime `json:"-" gorm:"index"`
}

func (Role) TableName() string {
	return "p_roles"
}

func (Menu) TableName() string {
	return "p_menus"
}
