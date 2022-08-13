package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exceptionHandler struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"status_code"`
	Service    string `json:"service"`
}

func HttpErrorHandler(errCode int, message string, c *gin.Context) {
	switch errCode {
	// 400
	case http.StatusBadRequest:
		c.JSON(http.StatusBadRequest, &exceptionHandler{
			Message:    fmt.Sprintf("Bad request, %v", message),
			StatusCode: http.StatusBadRequest,
		})
	// 403
	case http.StatusForbidden:
		c.JSON(http.StatusForbidden, &exceptionHandler{
			Message:    fmt.Sprintf("Forbidden, %v", message),
			StatusCode: http.StatusForbidden,
		})
	// 500
	default:
		c.JSON(http.StatusInternalServerError, &exceptionHandler{
			Message:    fmt.Sprintf("Unknown error, %v", message),
			StatusCode: http.StatusInternalServerError,
		})
	}
}
