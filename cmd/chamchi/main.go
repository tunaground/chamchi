package main

import (
	"github.com/tunarider/chamchi/internal/command"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:     "Chamchi",
		Usage:    "스레드 형식 웹 게시판 API 서버",
		Version:  "v0.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "참치라이더",
				Email: "tunarider@tunaground.net",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "데이터베이스 구성",
				Action:  command.Migrate,
				Flags: append(
					command.MigrateFlags(),
					command.DbFlags()...,
				),
			},
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Chamchi 서버 구동",
				Action:  command.Serve,
				Flags: append(
					command.ServeFlags(),
					command.DbFlags()...,
				),
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
