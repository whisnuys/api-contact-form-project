package handlers

import (
	"api-contact-form/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainHandler handles HTTP requests for the main/root endpoint.
type MainHandler struct{}

// NewMainHandler creates a new instance of MainHandler.
func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

// MainHandler responds with a message indicating that the API Contact Form is running.
//
// It returns a JSON response with a 200 OK status code and a success message.
//
// Example Response:
//
//	{
//	    "code": "SUCCESS",
//	    "message": "API Contact Form is running."
//	}
func (h *MainHandler) MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API Contact Form is running.",
	})
}
