package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	UsersInfo UsersInfo
}

type UsersInfo struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	UserId    int    `json:"userId" gorm:"foreignKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Sessions struct {
	Id        int    `gorm:"primaryKey"`
	SessionId string `json:"sessionId"`
	UserId    int    `json:"userId" gorm:"foreignKey"`
}
