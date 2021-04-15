package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/service"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

type CreateThreadInput struct {
	BoardID uint `json:"board_id"`
	Title string `json:"title"`
	Password string `json:"password"`
}

func CreateThread(ctx *context.Context) gin.HandlerFunc {
	gb := service.GetBoard(ctx)
	ct := service.CreateThread(ctx)
	return func(c *gin.Context) {
		var input CreateThreadInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board, count, err := gb(int(input.BoardID))
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread := model.Thread{
			BoardID: board.ID,
			Title: input.Title,
			Password: input.Password,
		}
		data, err := ct(thread)
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
	BoardID uint `json:"board_id"`
}

func GetThreads(ctx *context.Context) gin.HandlerFunc {
	gts := service.GetThreads(ctx)
	return func(c *gin.Context) {
		var input GetThreadsInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		var pagination util.Pagination
		if err := c.Bind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		threads, count, err := gts(int(input.BoardID), pagination.Offset, pagination.Limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
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
	gt := service.GetThread(ctx)
	return func(c *gin.Context) {
		thread, count, err := gt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
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
	BoardID uint `json:"board_id"`
	Title string `json:"title"`
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}

func UpdateThread(ctx *context.Context) gin.HandlerFunc {
	gb := service.GetBoard(ctx)
	gt := service.GetThread(ctx)
	ut := service.UpdateThread(ctx)
	return func(c *gin.Context) {
		var input UpdateThreadInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		thread, count, err := gt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
			return
		}
		if thread.Password != input.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong thread password"})
			return
		}
		if thread.BoardID != input.BoardID {
			board, count, err := gb(int(input.BoardID))
			if count == 0 {
				c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
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
		data, err := ut(thread)
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
