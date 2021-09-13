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

func GetThreads(ctx *context.Context) func(uint, model.ThreadStatus, int, int) ([]model.Thread, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(boardId uint, status model.ThreadStatus, offset int, limit int) (threads []model.Thread, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return threads, count, err
		}
		db.Where(&model.Thread{BoardID: boardId, Status: status}).Order("updated_at asc").Limit(limit).Offset(offset).Find(&threads).Count(&count)
		return threads, count, err
	}
}

func GetThread(ctx *context.Context) func(uint, model.ThreadStatus) (model.Thread, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(id uint, status model.ThreadStatus) (thread model.Thread, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return thread, count, err
		}
		if status == model.ThreadStatusAll {
			db.Where(&model.Thread{ID: id}).Find(&thread).Count(&count)
		} else {
			db.Where(&model.Thread{ID: id, Status: status}).Find(&thread).Count(&count)
		}
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
