package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title" gorm:"not null;size:255" binding:"required,min=1,max=255"`
	Author          string         `json:"author" gorm:"not null;size:255" binding:"required,min=1,max=255"`
	ISBN            string         `json:"isbn" gorm:"uniqueIndex;size:20" binding:"required,min=10,max=20"`
	PublishedYear   int            `json:"published_year" gorm:"not null" binding:"required,min=1000,max=2100"`
	Genre           string         `json:"genre" gorm:"size:100" binding:"required,min=1,max=100"`
	AvailableCopies int            `json:"available_copies" gorm:"not null;default:0" binding:"min=0"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateBookRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=255"`
	Author          string `json:"author" binding:"required,min=1,max=255"`
	ISBN            string `json:"isbn" binding:"required,min=10,max=20"`
	PublishedYear   int    `json:"published_year" binding:"required,min=1000,max=2100"`
	Genre           string `json:"genre" binding:"required,min=1,max=100"`
	AvailableCopies int    `json:"available_copies" binding:"min=0"`
}

type UpdateBookRequest struct {
	Title           *string `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Author          *string `json:"author,omitempty" binding:"omitempty,min=1,max=255"`
	ISBN            *string `json:"isbn,omitempty" binding:"omitempty,min=10,max=20"`
	PublishedYear   *int    `json:"published_year,omitempty" binding:"omitempty,min=1000,max=2100"`
	Genre           *string `json:"genre,omitempty" binding:"omitempty,min=1,max=100"`
	AvailableCopies *int    `json:"available_copies,omitempty" binding:"omitempty,min=0"`
}

func (r *CreateBookRequest) ToBook() *Book {
	return &Book{
		Title:           r.Title,
		Author:          r.Author,
		ISBN:            r.ISBN,
		PublishedYear:   r.PublishedYear,
		Genre:           r.Genre,
		AvailableCopies: r.AvailableCopies,
	}
}

func (r *UpdateBookRequest) UpdateBook(book *Book) {
	if r.Title != nil {
		book.Title = *r.Title
	}
	if r.Author != nil {
		book.Author = *r.Author
	}
	if r.ISBN != nil {
		book.ISBN = *r.ISBN
	}
	if r.PublishedYear != nil {
		book.PublishedYear = *r.PublishedYear
	}
	if r.Genre != nil {
		book.Genre = *r.Genre
	}
	if r.AvailableCopies != nil {
		book.AvailableCopies = *r.AvailableCopies
	}
}
