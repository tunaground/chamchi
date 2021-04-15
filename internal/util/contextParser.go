package util

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type ContextParser struct {
	Context *context.Context
}

func (c ContextParser) Database() (*gorm.DB, error) {
	ctx := *c.Context
	i := ctx.Value("db")
	db, ok := i.(*gorm.DB)
	if ok {
		return db, nil
	} else {
		return db, errors.New("type assertion failed")
	}
}
