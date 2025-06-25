package middleware

import(
	"log"
	"time"
	"github.com/gin-gonic/gin"
)


func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// After request
		duration := time.Since(start)
		log.Printf("[%s] %s - %d (%s)", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

