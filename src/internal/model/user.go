package model

type User struct {
	UserId   uint64 `json:"user_id" gorm:"autoIncrement;primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user_auth"
}
