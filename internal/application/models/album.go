package models

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title" binding:"required,lt=200"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

func NewAlbum(Id int64, title string, artist string, price float64) *Album {
	return &Album{ID: Id, Title: title, Artist: artist, Price: price}
}
