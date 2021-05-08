package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func  Echo(c *gin.Context) {
	queryString, ok := c.GetQuery("query")
	if !ok || len(strings.ReplaceAll(queryString, " ", "")) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "query not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"echo":queryString})
	return
}

func  GetTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time":time.Now()})
	return
}
