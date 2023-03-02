package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/chatgpt-go/internal/pkg/service"
	"net/http"
)

// Handler entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.POST("api/chat", service.Chat)
	router.ServeHTTP(w, r)
}
