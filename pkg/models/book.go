package models

import (
	"time"
)

type CreateBookParams struct {
	GenreID     uint32 `json:"genreID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorId    uint32
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
	ID          uint32    `json:"id" redis:"id"`
	Author      string    `json:"author" redis:"author"`
	Genre       Genre     `json:"genre" redis:"genre"`
	Title       string    `json:"title" redis:"title"`
	Description string    `json:"description" redis:"description"`
	ImgURL      string    `json:"imageURL" redis:"imageURL"`
	CreatedAt   time.Time `json:"createdAt" redis:"createdAt"`
}

type UpdateBookParams struct {
	GenreID     uint32 `json:"genreID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorId    uint32
	BookId      uint32
}

type BookFilter struct {
	AuthorId uint32
	GenreId  uint32
}

type SearchParams struct {
	SearchWord    string
	CreatedAtSort string
	Pagination    PaginationParams
	Filter        BookFilter
}

type BooksParams struct {
	Filter        BookFilter
	Pagination    PaginationParams
	CreatedAtSort string
}
