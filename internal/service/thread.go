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

func GetThreads(ctx *context.Context) func(int, model.ThreadStatus, int, int) ([]model.Thread, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(boardId int, status model.ThreadStatus, offset int, limit int) (threads []model.Thread, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return threads, count, err
		}
		db.Model(&model.Thread{}).Where("board_id = ? AND status = ?", boardId, status).Order("updated_at asc").Limit(limit).Offset(offset).Find(&threads).Count(&count)
		return threads, count, err
	}
}

func GetThread(ctx *context.Context) func(int, model.ThreadStatus) (model.Thread, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(id int, status model.ThreadStatus) (thread model.Thread, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return thread, count, err
		}
		db.Model(&model.Thread{}).Where("id = ? AND status = ?", id, status).Find(&thread).Count(&count)
		db.Model(&thread).Association("responses")
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
