package dtos

type Album struct {
	ID     string  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Year   int     `json:"year"`
}
