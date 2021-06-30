package db

type User struct {
	Count 		string 		`gorm:"primaryKey";json:"count"`
	PassWd		string 		`gorm:"primaryKey";json:"pass_wd"`
}
