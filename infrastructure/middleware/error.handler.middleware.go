package middleware

import (
	"eCommerce/errs"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			errs.HTTPErrorResponse(c.Writer, e.Err)
		}
	}
}
