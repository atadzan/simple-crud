package models

import "time"

type CreateBookParams struct {
	GenreID     uint32 `json:"genreID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Author struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

type BookList struct {
	ID        uint32    `json:"id"`
	Author    Author    `json:"author"`
	Genre     Genre     `json:"genre"`
	Title     string    `json:"title"`
	ImgURL    string    `json:"imageURL"`
	CreatedAt time.Time `json:"createdAt"`
}

type BookDetailed struct {
	ID          uint32    `json:"id"`
	Author      string    `json:"author"`
	Genre       Genre     `json:"genre"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImgURL      string    `json:"imageURL"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UpdateBookParams struct {
	GenreID     uint32 `json:"genreID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
