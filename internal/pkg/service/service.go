package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/chatgpt-go/pkg/utils/log"
	"github.com/otiai10/openaigo"
	"strconv"

	"github.com/lyleshaw/chatgpt-go/internal/pkg/openai"
)

type ChatContext struct {
	ConversationId  string `json:"conversationId"`
	ParentMessageId string `json:"parentMessageId"`
}

type ChatRequest struct {
	Prompt  string       `json:"prompt"`
	Options *ChatContext `json:"options"`
}

type ChatResponse struct {
	Role            string `json:"role"`
	ID              string `json:"id"`
	ParentMessageId string `json:"parentMessageId"`
	ConversationId  string `json:"conversationId"`
	Text            string `json:"text"`
}

type Response struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    ChatResponse `json:"data"`
}

var SessionList = make(map[string][]openaigo.ChatMessage)
var Count = 0

// Index .
// @router / [GET]
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// Chat .
// @router /chat [POST]
func Chat(c *gin.Context) {
	var req ChatRequest
	var response Response
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Infof("req: %v", req)

	var message []openaigo.ChatMessage
	var id string
	if req.Options == nil {
		message = append(message, openaigo.ChatMessage{
			Role:    "user",
			Content: req.Prompt,
		})
		SessionList[strconv.Itoa(Count)] = message
		id = strconv.Itoa(Count)
		Count++
	} else {
		message = append(message, SessionList[req.Options.ParentMessageId]...)
		message = append(message, openaigo.ChatMessage{
			Role:    "user",
			Content: req.Prompt,
		})
		SessionList[req.Options.ParentMessageId] = message
		id = strconv.Itoa(Count)
	}

	resp, err := openai.Chat(message)
	response.Data.Text = resp

	if err != nil {
		response.Status = "Fail"
		response.Message = "Something went wrong"
		c.JSON(400, response)
		return
	}
	response.Status = "Success"
	response.Message = ""
	c.JSON(200, Response{
		Status:  "Success",
		Message: "",
		Data: ChatResponse{
			Role:            "assistant",
			ID:              id,
			ParentMessageId: id,
			ConversationId:  id,
			Text:            resp,
		},
	})
	return
}
