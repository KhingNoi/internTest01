package models

type User struct {
	UserId    int    `gorm:"primaryKey"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
	AvatarURL string `json:"AvatarURL"`
}
