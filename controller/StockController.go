package controller

import (
	repos "Main/repository"
)

type BaseHandler struct {
	stockRepo repos.StockRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(stockRepo repos.StockRepository) *BaseHandler {
	return &BaseHandler{
		stockRepo: stockRepo,
	}
}
