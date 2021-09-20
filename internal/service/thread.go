package service

import (
	"context"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

func CreateThread(ctx *context.Context) func(model.Thread) (model.Thread, error) {
	cp := util.ContextParser{Context: ctx}
	return func(thread model.Thread) (model.Thread, error) {
		db, err := cp.Database()
		if err != nil {
			return thread, err
		}
		if result := db.Create(&thread); result.Error != nil {
			return thread, result.Error
		}
		return thread, nil
	}
}

func GetThreads(ctx *context.Context) func(map[string]interface{}, int, int) ([]model.Thread, int, error) {
	cp := util.ContextParser{Context: ctx}
	return func(query map[string]interface{}, offset int, limit int) (threads []model.Thread, count int, err error) {
		db, err := cp.Database()
		if err != nil {
			return threads, count, err
		}
		db.Where(query).Order("updated_at desc").Limit(limit).Offset(offset).Find(&threads)
		return threads, len(threads), err
	}
}

func GetThread(ctx *context.Context) func(map[string]interface{}, model.ThreadStatus) (model.Thread, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(query map[string]interface{}, status model.ThreadStatus) (thread model.Thread, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return thread, count, err
		}
		if status != model.ThreadStatusAll {
			query["status"] = status
		}
		db.Where(query).Find(&thread).Count(&count)
		err = db.Model(&thread).Association("Responses").Find(&thread.Responses)
		if err != nil {
			return thread, count, err
		}
		return thread, count, nil
	}
}

func UpdateThread(ctx *context.Context) func(model.Thread) (model.Thread, error) {
	cp := util.ContextParser{Context: ctx}
	return func(thread model.Thread) (model.Thread, error) {
		db, err := cp.Database()
		if err != nil {
			return thread, err
		}
		db.Save(&thread)
		return thread, nil
	}
}
