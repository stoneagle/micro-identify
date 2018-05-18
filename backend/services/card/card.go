package services

import "github.com/go-xorm/xorm"

type Card struct {
	engine *xorm.Engine
}

func NewCard(engine *xorm.Engine) *Card {
	return &Card{
		engine: engine,
	}
}
