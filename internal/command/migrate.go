package command

import (
	"errors"
	"fmt"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MigrateFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "기존에 존재하는 데이터 삭제 후 재생성",
		},
	}
}

func Migrate(c *cli.Context) error {
	dsn := c.String("dsn")
	force := c.Bool("force")
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
		return cli.Exit("connect to database failed", 0)
	}
	err = createTable(db, force, model.Board{})
	if err != nil {
		return cli.Exit(util.StackError(errors.New("failed to create board table"), err), 1)
	}
	err = createTable(db, force, model.Thread{})
	if err != nil {
		return cli.Exit(util.StackError(errors.New("failed to create thread table"), err), 1)
	}
	return nil
}

func createTable(db *gorm.DB, force bool, obj interface{}) error {
	if !db.Migrator().HasTable(obj) {
		err := db.Migrator().CreateTable(obj)
		if err != nil {
			return cli.Exit(err, 1)
		}
	} else {
		if force {
			err := db.Migrator().DropTable(obj)
			if err != nil {
				return cli.Exit(err, 2)
			}
			err = db.Migrator().CreateTable(obj)
			if err != nil {
				return cli.Exit(err, 2)
			}
		} else {
			return cli.Exit("table already exists", 3)
		}
	}
	return nil
}
