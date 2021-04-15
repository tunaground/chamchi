package command

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tunarider/chamchi/internal/middleware"
	"github.com/tunarider/chamchi/internal/route"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func ServeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "listen",
			Aliases: []string{"l"},
			Usage:   "서버 IP",
			Value:   "127.0.0.1",
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "서버 포트",
			Value:   8080,
		},
	}
}

func Serve(c *cli.Context) error {
	ctx := context.Background()

	dsn := c.String("dsn")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.String("user"),
			c.String("password"),
			c.String("host"),
			c.String("database"),
		)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return cli.Exit("failed to connect to database", 1)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return cli.Exit("failed to connect to database", 1)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	ctx = context.WithValue(ctx, "db", db)

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Options([]string{"OPTIONS", "POST", "GET", "PUT"}))
	route.Route(&ctx, engine.Group("/"))
	err = engine.Run(fmt.Sprintf("%s:%d", c.String("host"), c.Int("port")))
	if err != nil {
		return cli.Exit("failed to start engine", 1)
	}
	return nil
}
