package vo

// Address houses a users address information
type Address struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	Planet string `json:"planet" binding:"required"`
	Phone  string `json:"phone" binding:"required"`
}
type User struct {
	FirstName      string    `json:"firstName" binding:"required"`
	LastName       string    `json:"lastName" binding:"required"`
	Age            uint8     `json:"age" binding:"gte=0,lte=130"`
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" form:"password" binding:"required"`
	Mobile         uint      `json:"mobile"`
	FavouriteColor string    `json:"favouriteColor" binding:"iscolor"`           // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []Address `json:"addresses" binding:"required,dive,required"` // a person can have a home and cottage...
}

type Policy struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
	Obj string `json:"obj" form:"obj" binding:"required"`
	Act string `json:"act" form:"act" binding:"required"`
}

type LoginForm struct {
	Password string `json:"password" form:"password" binding:"required"`
	Mobile   uint   `json:"mobile" form:"mobile" binding:"required"`
}
