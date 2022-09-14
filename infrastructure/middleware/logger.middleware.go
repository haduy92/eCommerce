package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := GetDurationInMillSeconds(start)

		if c.Writer.Status() < 500 {
			log.WithFields(log.Fields{
				"duration":   duration,
				"method":     c.Request.Method,
				"path":       c.Request.RequestURI,
				"status":     c.Writer.Status(),
				"referrer":   c.Request.Referer(),
				"request_id": c.Writer.Header().Get(CorrelationHeader),
			}).Info("")
		}
	}
}

func GetDurationInMillSeconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
