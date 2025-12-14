package validator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(err error, c *gin.Context) {
	if err != nil {
		var validationErrors = err.(validator.ValidationErrors)
		var errors = make(map[string]string)
		
		for _, e := range validationErrors {
			switch(e.Tag()) {
				case "required":
					errors[e.Field()] = "is required"
				case "gt":
					errors[e.Field()] = "must be greater than " + e.Param()
			}
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "fields": errors})
		return
	}
}
