package service

import (
	"context"
	"errors"
	"github.com/tunarider/chamchi/internal/util"
	"github.com/tunarider/chamchi/pkg/model"
)

func CreateBoard(ctx *context.Context) func(model.Board) (model.Board, error) {
	cp := util.ContextParser{Context: ctx}
	return func(board model.Board) (model.Board, error) {
		db, err := cp.Database()
		if err != nil {
			return board, err
		}
		var count int64
		db.Model(&model.Board{}).Where("name = ?", board.Name).Count(&count)
		if count != 0 {
			return board, errors.New("board already exists")
		}
		if result := db.Create(&board); result.Error != nil {
			return board, result.Error
		}
		return board, nil
	}
}

func GetBoards(ctx *context.Context) func(int, int) ([]model.Board, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(offset int, limit int) (boards []model.Board, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return boards, count, err
		}
		db.Order("name asc").Limit(limit).Offset(offset).Find(&boards).Count(&count)
		return boards, count, err
	}
}

func GetBoard(ctx *context.Context) func(int) (model.Board, int64, error) {
	cp := util.ContextParser{Context: ctx}
	return func(id int) (board model.Board, count int64, err error) {
		db, err := cp.Database()
		if err != nil {
			return board, count, err
		}
		db.Model(&model.Board{}).Where("id = ?", id).Find(&board).Count(&count)
		return board, count, nil
	}
}

func UpdateBoard(ctx *context.Context) func(model.Board) (model.Board, error) {
	cp := util.ContextParser{Context: ctx}
	return func(board model.Board) (model.Board, error) {
		db, err := cp.Database()
		if err != nil {
			return board, err
		}
		db.Save(&board)
		return board, nil
	}
}
