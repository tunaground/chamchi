package service

import (
	"context"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

func CreateResponse(ctx *context.Context) func(response model.Response) (model.Response, error) {
	cp := util.ContextParser{Context: ctx}
	return func(response model.Response) (model.Response, error) {
		db, err := cp.Database()
		if err != nil {
			return response, err
		}
		if result := db.Create(&response); result.Error != nil {
			return response, result.Error
		}
		return response, nil
	}
}

func GetResponses(ctx *context.Context) func(int, int, int) ([]model.Response, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(threadId int, offset int, limit int) (response []model.Response, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return response, count, err
		}
		db.Model(&model.Response{}).Where("thread_id = ?", threadId).Order("sequence asc").Limit(limit).Offset(offset).Find(&response).Count(&count)
		return response, count, nil
	}
}

func GetResponse(ctx *context.Context) func(int) (model.Response, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(id int) (response model.Response, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return response, count, err
		}
		db.Model(&model.Response{}).Where("id = ?", id).Find(&response).Count(&count)
		return response, count, nil
	}
}

func UpdateResponse(ctx *context.Context) func(model.Response) (model.Response, error) {
	cp := util.ContextParser{Context: ctx}
	return func(response model.Response) (model.Response, error) {
		db, err := cp.Database()
		if err != nil {
			return response, err
		}
		db.Save(&response)
		return response, nil
	}
}
