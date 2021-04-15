package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/controller"
)

func SetRouteRoot(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("/health", controller.Health())

	r.GET("/boards", controller.GetBoards(ctx))
	r.GET("/boards/:id", controller.GetBoard(ctx))
	r.POST("/boards", controller.CreateBoard(ctx))
	r.PUT("/boards/:id", controller.UpdateBoard(ctx))
}
