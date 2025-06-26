package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   errorMessage,
	})
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, err error) {
	var errorMessages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, e.Field()+" is required")
			case "min":
				errorMessages = append(errorMessages, e.Field()+" must be at least "+e.Param()+" characters long")
			case "max":
				errorMessages = append(errorMessages, e.Field()+" must be at most "+e.Param()+" characters long")
			case "email":
				errorMessages = append(errorMessages, e.Field()+" must be a valid email")
			default:
				errorMessages = append(errorMessages, e.Field()+" is invalid")
			}
		}
	} else {
		errorMessages = append(errorMessages, err.Error())
	}

	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: "Validation failed",
		Error:   strings.Join(errorMessages, ", "),
	})
}
