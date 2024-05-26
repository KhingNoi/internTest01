package models

type User struct {
	Id        int     `gorm:"primaryKey" json:"id"`
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	AvatarURL string  `json:"avatar_url"`
	Orders    []Order `gorm:"foreignKey:UserID"`
}
