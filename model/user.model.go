package model

// Address houses a users address information
type User struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"firstName" gorm:"not null;type:varchar(50);"`
	LastName       string `json:"lastName" gorm:"not null;type:varchar(50);"`
	Age            uint8  `json:"age" gorm:"not null;type:tinyint;"`
	Email          string `json:"email" gorm:"type:varchar(50);"`
	Mobile         uint   `json:"mobile" gorm:"type:varchar(11);uniqueIndex;not null;"`
	Password       string `json:"-" gorm:"type:varchar(255);not null;"`
	FavouriteColor string `json:"favouriteColor,omitempty"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []UserAddress
	Roles          []*Role  `gorm:"many2many:p_user_roles;"`
	CreatedAt      NullTime `json:"createdAt,omitempty"`
	UpdatedAt      NullTime `json:"updatedAt,omitempty"`
	DeletedAt      NullTime `json:"deletedAt,omitempty" gorm:"index"`
}

// Address houses a users address information
type UserAddress struct {
	ID        uint     `json:"id"`
	UserID    uint     `json:"userId"`
	User      User     `json:"-"`
	Street    string   `json:"street" gorm:"not null;type:varchar(50);"`
	City      string   `json:"city" gorm:"not null;type:varchar(50);"`
	Planet    string   `json:"planet" gorm:"not null;type:varchar(50);"`
	Phone     string   `json:"phone" gorm:"not null;type:varchar(15);"`
	CreatedAt NullTime `json:"createdAt"`
	UpdatedAt NullTime `json:"updatedAt"`
	DeletedAt NullTime `json:"deletedAt" gorm:"index"`
}

func (User) TableName() string {
	return "m_users"
}

func (UserAddress) TableName() string {
	return "m_user_address"
}
