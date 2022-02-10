package controllers

import (
	repositories "Main/Repositories"
)

type BaseHandler struct {
	stockRepo repositories.StockRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(stockRepo repositories.StockRepository) *BaseHandler {
	return &BaseHandler{
		stockRepo: stockRepo,
	}
}
