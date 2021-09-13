package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/service"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

type CreateThreadInput struct {
	BoardID  uint   `json:"board_id"`
	Title    string `json:"title"`
	Password string `json:"password"`
}

func CreateThread(ctx *context.Context) gin.HandlerFunc {
	getBoardService := service.GetBoard(ctx)
	createThreadService := service.CreateThread(ctx)
	return func(c *gin.Context) {
		var input CreateThreadInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board, count, err := getBoardService(int(input.BoardID))
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread := model.Thread{
			BoardID:  board.ID,
			Title:    input.Title,
			Password: input.Password,
			Status:   model.ThreadStatusPrepare,
		}
		data, err := createThreadService(thread)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data": gin.H{
				"thread": data,
			},
		})
	}
}

type GetThreadsInput struct {
	BoardID uint `form:"board_id"`
}

func GetThreads(ctx *context.Context) gin.HandlerFunc {
	getThreadsService := service.GetThreads(ctx)
	return func(c *gin.Context) {
		var input GetThreadsInput
		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		var pagination util.Pagination
		if err := c.Bind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		threads, count, err := getThreadsService(input.BoardID, model.ThreadStatusConfirm, pagination.Offset, pagination.Limit)
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
				"threads": threads,
			},
		})
	}
}

func GetThread(ctx *context.Context) gin.HandlerFunc {
	getThreadService := service.GetThread(ctx)
	return func(c *gin.Context) {
		threadId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread, count, err := getThreadService(uint(threadId), model.ThreadStatusConfirm)
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
				"thread": thread,
			},
		})
	}
}

type UpdateThreadInput struct {
	BoardID     uint   `json:"board_id"`
	Title       string `json:"title"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type RouteUpdateThreadQuery struct {
	Confirm string `form:"confirm"`
}

func RouteUpdateThread(ctx *context.Context) gin.HandlerFunc {
	ct := confirmThread(ctx)
	ut := updateThread(ctx)
	return func(c *gin.Context) {
		var query RouteUpdateThreadQuery
		if err := c.Bind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if query.Confirm == "true" {
			ct(c)
		} else {
			ut(c)
		}
	}
}

func confirmThread(ctx *context.Context) gin.HandlerFunc {
	getThreadService := service.GetThread(ctx)
	updateThreadService := service.UpdateThread(ctx)
	return func(c *gin.Context) {
		threadId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread, count, err := getThreadService(uint(threadId), model.ThreadStatusPrepare)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		if len(thread.Responses) == 1 {
			thread.Status = model.ThreadStatusConfirm
			data, err := updateThreadService(thread)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"message": "ok",
				"data": gin.H{
					"thread": data,
				},
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "thread is empty"})
			return
		}
	}
}

func updateThread(ctx *context.Context) gin.HandlerFunc {
	getBoardService := service.GetBoard(ctx)
	getThreadService := service.GetThread(ctx)
	updateThreadService := service.UpdateThread(ctx)
	return func(c *gin.Context) {
		var input UpdateThreadInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		threadId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread, count, err := getThreadService(uint(threadId), model.ThreadStatusConfirm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		if thread.Password != input.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong thread password"})
			return
		}
		if thread.BoardID != input.BoardID {
			board, count, err := getBoardService(int(input.BoardID))
			if count == 0 {
				c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
				return
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
			thread.BoardID = board.ID
		}
		thread.BoardID = input.BoardID
		thread.Title = input.Title
		if input.NewPassword != "" {
			thread.Password = input.NewPassword
		}
		data, err := updateThreadService(thread)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data": gin.H{
				"thread": data,
			},
		})
	}
}
