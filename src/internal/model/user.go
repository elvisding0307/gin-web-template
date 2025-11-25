package model

type User struct {
	UserId     uint64 `json:"user_id" gorm:"autoIncrement;primaryKey"`
	Username   string `json:"username" gorm:"primaryKey"`
	Password   string `json:"password"`
	Email      string `json:"email" gorm:"primaryKey"`
	Phone      string `json:"phone" gorm:"primaryKey"`
	Permission string `json:"permission" gorm:"default:'user'"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
