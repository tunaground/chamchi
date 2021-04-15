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

type CreateBoardInput struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func CreateBoard(ctx *context.Context) gin.HandlerFunc {
	cb := service.CreateBoard(ctx)
	return func(c *gin.Context) {
		var input CreateBoardInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board := model.Board{Key: input.Key, Name: input.Name}
		data, err := cb(board)
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
	gbs := service.GetBoards(ctx)
	return func(c *gin.Context) {
		var pagination util.Pagination
		if err := c.Bind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		boards, count, err := gbs(pagination.Offset, pagination.Limit)
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
				"boards": boards,
			},
		})
	}
}

func GetBoard(ctx *context.Context) gin.HandlerFunc {
	gb := service.GetBoard(ctx)
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board, count, err := gb(id)
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
				"board": board,
			},
		})
	}
}

type UpdateBoardInput struct {
	Name string `json:"Name"`
}

func UpdateBoard(ctx *context.Context) gin.HandlerFunc {
	gb := service.GetBoard(ctx)
	ub := service.UpdateBoard(ctx)
	return func(c *gin.Context) {
		var input UpdateBoardInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		board, count, err := gb(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "ok"})
			return
		}
		board.Name = input.Name
		data, err := ub(board)
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
