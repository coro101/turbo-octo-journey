package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!!!!",
			"dog":     "lei a",
			"person":  "hyowon",
			"cup":     "ramyeon",
			"height":  160,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
