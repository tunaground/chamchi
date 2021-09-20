package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/service"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
	"net/http"
)

func getUserIP(c *gin.Context) string {
	return c.ClientIP()
}

func createUserId(c *gin.Context) string {
	return getUserIP(c)[:1]
}

type CreateResponseInput struct {
	ThreadID   uint   `json:"thread_id"`
	Username   string `json:"username"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
	Youtube    string `json:"youtube"`
}

func CreateResponse(ctx *context.Context) gin.HandlerFunc {
	getThreadService := service.GetThread(ctx)
	createResponseService := service.CreateResponse(ctx)
	return func(c *gin.Context) {
		var input CreateResponseInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread, count, err := getThreadService(
			map[string]interface{}{"id": input.ThreadID},
			model.ThreadStatusAll,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		var seq uint
		if len(thread.Responses) == 0 {
			seq = 0
		} else {
			seq = thread.Responses[len(thread.Responses)-1].Sequence
		}
		response := model.Response{
			ThreadID:   thread.ID,
			Sequence:   seq,
			Username:   input.Username,
			UserID:     createUserId(c),
			IP:         getUserIP(c),
			Content:    input.Content,
			Attachment: input.Attachment,
			Youtube:    input.Youtube,
		}
		data, err := createResponseService(response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data": gin.H{
				"response": data,
			},
		})
	}
}

func GetResponses(ctx *context.Context) gin.HandlerFunc {
	getResponsesService := service.GetResponses(ctx)
	return func(c *gin.Context) {
		query := map[string]interface{}{}
		if threadID, ok := c.GetQuery("thread_id"); ok {
			query["thread_id"] = threadID
		}
		if sequence, ok := c.GetQuery("sequence"); ok {
			query["sequence"] = sequence
		}

		var pagination util.Pagination
		if err := c.Bind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		responses, count, err := getResponsesService(query, pagination.Offset, pagination.Limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data": gin.H{
				"responses": responses,
			},
		})
	}
}
