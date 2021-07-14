package handlers

import "github.com/gin-gonic/gin"

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "{"c": "pong"}"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"c": "pong",
	})
}
