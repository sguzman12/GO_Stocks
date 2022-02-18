package service

import (
	entity "Main/entity"
)

var (
	StockService stockServiceInterface = &stockService{}
)

type stockServiceInterface interface {
	Create(entity.Stock) *entity.Stock
	GetById(entity.Stock) *entity.Stock
	// GetAll([]entity.Stock) *entity.Stock
}

type stockService struct {
}

func (s *stockService) Create(entity.Stock) *entity.Stock {
	return nil
}

func (s *stockService) GetById(entity.Stock) *entity.Stock {
	return nil
}

// func (s *stockService) GetAll(entity.Stock) *entity.Stock {

// 	// stockArray := []entity.Stock{}
// 	// return &stockArray
// 	return nil
// }
