package services

import (
	"context"
	"fmt"
	"strconv"

	"library-management-system/internal/models"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		db: db,
	}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book

	//get from database
	if err := s.db.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch books from database: %w", err)
	}

	return books, nil
}

func (s *BookService) GetBookByID(ctx context.Context, id uint) (*models.Book, error) {
	var book models.Book

	//get from database
	if err := s.db.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to fetch book from database: %w", err)
	}

	return &book, nil
}

func (s *BookService) CreateBook(ctx context.Context, req *models.CreateBookRequest) (*models.Book, error) {
	book := req.ToBook()

	// Check if ISBN already exists
	var existingBook models.Book
	if err := s.db.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err == nil {
		return nil, fmt.Errorf("book with ISBN %s already exists", book.ISBN)
	}

	if err := s.db.Create(book).Error; err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return book, nil
}

func (s *BookService) UpdateBook(ctx context.Context, id uint, req *models.UpdateBookRequest) (*models.Book, error) {
	// First, get the existing book
	book, err := s.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if ISBN is being updated and if it conflicts with another book
	if req.ISBN != nil && *req.ISBN != book.ISBN {
		var existingBook models.Book
		if err := s.db.Where("isbn = ? AND id != ?", *req.ISBN, id).First(&existingBook).Error; err == nil {
			return nil, fmt.Errorf("book with ISBN %s already exists", *req.ISBN)
		}
	}

	// Update the book with new values
	req.UpdateBook(book)

	if err := s.db.Save(book).Error; err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return book, nil
}

func (s *BookService) DeleteBook(ctx context.Context, id uint) error {
	// Check if book exists
	_, err := s.GetBookByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&models.Book{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

func (s *BookService) ParseID(idStr string) (uint, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid book ID")
	}
	return uint(id), nil
}
