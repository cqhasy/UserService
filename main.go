package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {})
	r.POST("/register", func(c *gin.Context) {})
	r.PUT("/update_user", func(c *gin.Context) {})
	r.DELETE("/delete_user", func(c *gin.Context) {})
}
