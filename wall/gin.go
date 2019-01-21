package wall

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Paywall(c *gin.Context) {
	// Main Middleware for Gin Program
	log.Println("Validating the quota")
	conf.Parse()
	if HasQuota(c.GetHeader("UserIdentifier"), c.Request.URL) {
		c.Header("Paywall@AutoAI", "Authorized")
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "failed",
			"hint":   "The request is rejected due to insufficient quota",
			"url":    "",
			"code":   10403,
		})
	}
}
