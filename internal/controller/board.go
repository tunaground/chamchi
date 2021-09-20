package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/service"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

type CreateBoardInput struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func CreateBoard(ctx *context.Context) gin.HandlerFunc {
	createBoardService := service.CreateBoard(ctx)
	return func(c *gin.Context) {
		var input CreateBoardInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board := model.Board{Key: input.Key, Name: input.Name}
		data, err := createBoardService(board)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data": gin.H{
				"board": data,
			},
		})
	}
}

func GetBoards(ctx *context.Context) gin.HandlerFunc {
	getBoardsService := service.GetBoards(ctx)
	return func(c *gin.Context) {
		var pagination util.Pagination
		if err := c.Bind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		query := map[string]interface{}{}
		if key, ok := c.GetQuery("key"); ok {
			query["key"] = key
		}
		boards, count, err := getBoardsService(query, pagination.Offset, pagination.Limit)
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
				"boards": boards,
			},
		})
	}
}

type UpdateBoardInput struct {
	Name string `json:"name"`
}

func UpdateBoard(ctx *context.Context) gin.HandlerFunc {
	getBoardService := service.GetBoard(ctx)
	updateBoardService := service.UpdateBoard(ctx)
	return func(c *gin.Context) {
		query := map[string]interface{}{}
		if id, ok := c.GetQuery("id"); ok {
			query["id"] = id
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "empty id"})
			return
		}
		var input UpdateBoardInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board, count, err := getBoardService(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		board.Name = input.Name
		data, err := updateBoardService(board)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data": gin.H{
				"board": data,
			},
		})
	}
}
