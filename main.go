package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RepositoryRequest struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		var repositoryRequest RepositoryRequest
		if err := c.ShouldBindJSON(&repositoryRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"repositoryName": repositoryRequest.Name,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
