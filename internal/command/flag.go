package command

import "github.com/urfave/cli/v2"

func DbFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "dsn",
			Aliases: []string{"C"},
			Usage:   "데이터베이스 접속 정보. 제공시 다른 접속 정보 옵션 무시",
		},
		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"U"},
			Usage:   "데이터베이스 사용자 계정",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"P"},
			Usage:   "데이터베이스 사용자 암호",
		},
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"H"},
			Usage:   "데이터베이스 호스트",
		},
		&cli.StringFlag{
			Name:    "database",
			Aliases: []string{"D"},
			Usage:   "데이터베이스 이름",
		},
	}
}
