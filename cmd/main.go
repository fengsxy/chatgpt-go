package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/chatgpt-go/internal/pkg/service"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("public/index.html")
	router.GET("/", service.Index)
	router.POST("/api/chat", service.Chat)
	err := router.Run(":3003")
	if err != nil {
		return
	}
	return
}
