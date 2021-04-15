package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/controller"
)

func Route(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("/health", controller.Health())
	api := r.Group("/api")
	routeApi(ctx, api)
}

func routeApi(ctx *context.Context, r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	routeV1(ctx, v1)
}

func routeV1(ctx *context.Context, r *gin.RouterGroup) {
	boards := r.Group("/boards")
	routeBoards(ctx, boards)
	threads := r.Group("/threads")
	routeThreads(ctx, threads)
}

func routeBoards(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("", controller.GetBoards(ctx))
	r.GET("/:id", controller.GetBoard(ctx))
	r.POST("", controller.CreateBoard(ctx))
	r.PUT("/:id", controller.UpdateBoard(ctx))
}

func routeThreads(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("", controller.GetThreads(ctx))
	r.GET("/:id", controller.GetThread(ctx))
	r.POST("", controller.CreateThread(ctx))
	r.PUT("/:id", controller.RouteUpdateThread(ctx))
}
