package handlers

import (
	"net/http"

	"library-management-system/internal/models"
	"library-management-system/internal/services"
	"library-management-system/pkg/utils"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// GetBooks handles GET /api/v1/books
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch books", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Books retrieved successfully", books)
}

// GetBook handles GET /api/v1/books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := h.bookService.ParseID(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	book, err := h.bookService.GetBookByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "book not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "Book not found", err)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch book", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book retrieved successfully", book)
}

// CreateBook handles POST /api/v1/books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	book, err := h.bookService.CreateBook(c.Request.Context(), &req)
	if err != nil {
		if err.Error() == "book with ISBN "+req.ISBN+" already exists" {
			utils.ErrorResponse(c, http.StatusConflict, "Book with this ISBN already exists", err)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create book", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Book created successfully", book)
}

// UpdateBook handles PUT /api/v1/books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := h.bookService.ParseID(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	book, err := h.bookService.UpdateBook(c.Request.Context(), id, &req)
	if err != nil {
		if err.Error() == "book not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "Book not found", err)
			return
		}
		if err.Error() == "book with ISBN "+*req.ISBN+" already exists" {
			utils.ErrorResponse(c, http.StatusConflict, "Book with this ISBN already exists", err)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book updated successfully", book)
}

// DeleteBook handles DELETE /api/v1/books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := h.bookService.ParseID(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	if err := h.bookService.DeleteBook(c.Request.Context(), id); err != nil {
		if err.Error() == "book not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "Book not found", err)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete book", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book deleted successfully", nil)
}
