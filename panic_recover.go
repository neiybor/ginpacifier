package ginpacifier

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PanicRecovery is a Gin middleware that recovers from panics and returns a 500 error.
func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PanicRecovery] Recovered from panic: %v", err)
				if !c.Writer.Written() {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "Internal Server Error",
					})
				}
			}
		}()
		c.Next()
	}
}
