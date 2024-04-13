package entity

type User struct {
	Id       int64  `json:"id" gorm:"column:id; primary_key auto_increment"`
	Name     string `json:"name" gorm:"column:name"`
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "user"
}
