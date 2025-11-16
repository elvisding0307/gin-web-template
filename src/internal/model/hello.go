package model

type Hello struct {
	Id        uint64 `json:"id" gorm:"autoIncrement;primaryKey"`
	HelloWord string `json:"hello_word"`
}

// TableName 指定表名
func (Hello) TableName() string {
	return "hello"
}
