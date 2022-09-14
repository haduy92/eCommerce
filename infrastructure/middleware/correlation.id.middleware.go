package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// CorrelationHeader defines a default Correlation ID HTTP header.
const (
	CorrelationHeader = "X-Correlation-ID"
)

// SetRequestUUID will search for a correlation header and set a request-level
// correlation ID into the net.Context. If no header is found, a new UUID will
// be generated.
func SetCorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.Request.Header.Get(CorrelationHeader)
		if u == "" {
			u = uuid.NewV4().String()
		}
		c.Writer.Header().Set(CorrelationHeader, u)
		c.Next()
	}
}
