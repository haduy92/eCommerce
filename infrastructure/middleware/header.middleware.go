package middleware

import "github.com/gin-gonic/gin"

func SetJSONContentTypeHeader() gin.HandlerFunc {
	const (
		ContentTypeHeaderKey         string = "Content-Type"
		ContentTypeOptionsKey        string = "X-Content-Type-Options"
		JsonContentTypeHeaderVal     string = "application/json"
		NoSniffContentTypeOptionsVal string = "nosniff"
	)
	return func(c *gin.Context) {
		c.Writer.Header().Set(ContentTypeHeaderKey, JsonContentTypeHeaderVal)
		c.Writer.Header().Set(ContentTypeOptionsKey, NoSniffContentTypeOptionsVal)
		c.Next()
	}
}
