package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/controller"
)

func Route(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("/health", controller.Health())
	routeApi(ctx, r.Group("/api"))
}

func routeApi(ctx *context.Context, r *gin.RouterGroup) {
	routeV1(ctx, r.Group("/v1"))
}

func routeV1(ctx *context.Context, r *gin.RouterGroup) {
	routeBoards(ctx, r.Group("/board"))
	routeThreads(ctx, r.Group("/thread"))
	routeResponses(ctx, r.Group("/response"))
}

func routeBoards(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("", controller.GetBoards(ctx))
	r.POST("", controller.CreateBoard(ctx))
	r.PUT("", controller.UpdateBoard(ctx))
}

func routeThreads(ctx *context.Context, r *gin.RouterGroup) {
	r.GET("", controller.GetThreads(ctx))
	r.POST("", controller.CreateThread(ctx))
	r.PUT("", controller.RouteUpdateThread(ctx))
}

func routeResponses(ctx *context.Context, r *gin.RouterGroup) {
	r.POST("", controller.CreateResponse(ctx))
	r.GET("", controller.GetResponses(ctx))
}
