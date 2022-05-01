package models

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username" json:"username"`
	Password string `json:"password" json:"password"`
	State    string `json:"state" json:"state"`
	Age      int    `json:"age" json:"age"`
}
