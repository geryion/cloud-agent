package db

type User struct {
	Count 		string 		`gorm:"primaryKey";json:"count"`
	PassWd		string 		`gorm:"primaryKey";json:"pass_wd"`
}

type Music struct {
	Name 		string  	`gorm:"primaryKey";json:"name"`
	Time    	string      `json:"time"`
	Singer 		string    	`json:"singer"`

}
