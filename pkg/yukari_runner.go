package yukari

import (
	"yukari/pkg/database"
)

type Yukari struct {
	Store database.YukariDatabase
}

func NewYukari(store database.YukariDatabase) *Yukari {
	return &Yukari{Store: store}
}

func (y Yukari) Announce() error {
	return nil
}
