package model

type Hello struct {
	Id      uint64 `json:"id" gorm:"autoIncrement;primaryKey"`
	Content string `json:"content"`
}

// TableName 指定表名
func (Hello) TableName() string {
	return "hello"
}
