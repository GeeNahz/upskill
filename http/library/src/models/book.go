package models

type Status string

const (
	Read    Status = "read"
	Reading Status = "reading"
	ToRead  Status = "to_read"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status Status `json:"status" gorm:"default:to_read"` // read, reading, to_read
	Year   int    `json:"year"`
	UserID int    `json:"user_id"`
}
