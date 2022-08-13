package service

import (
	"github.com/gin-gonic/gin"
	"taoyuan_carpark/exception"
)

// Recovery if unexpected error cause that we catch error to response to user
func Recovery() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				exception.HttpErrorHandler(500, err.(string), c)
			}
		}()
		c.Next()
	}
}
