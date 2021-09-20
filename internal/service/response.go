package service

import (
	"context"
	"errors"
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
		count := int64(0)
		db.Where(&model.Response{ThreadID: response.ThreadID, Sequence: response.Sequence}).Count(&count)
		if count > 0 {
			return response, errors.New("duplicate response")
		}
		if result := db.Create(&response); result.Error != nil {
			return response, result.Error
		}
		return response, nil
	}
}

func GetResponses(ctx *context.Context) func(map[string]interface{}, int, int) ([]model.Response, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(query map[string]interface{}, offset int, limit int) (response []model.Response, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return response, count, err
		}
		db.Where(query).Order("sequence asc").Limit(limit).Offset(offset).Find(&response).Count(&count)
		return response, count, nil
	}
}

func GetResponse(ctx *context.Context) func(*model.Response) (model.Response, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(query *model.Response) (response model.Response, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return response, count, err
		}
		db.Where(query).Find(&response).Count(&count)
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
