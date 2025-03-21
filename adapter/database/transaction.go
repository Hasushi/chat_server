package database

import (
	"chat_server/usecase/output_port"

	"gorm.io/gorm"
)

type GormTransaction struct{
	db *gorm.DB
}

func NewGormTransaction(db *gorm.DB) output_port.GormTransaction {
	return &GormTransaction{db: db}
}

func (g *GormTransaction) StartTransaction(f func(tx interface{}) error) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		return f(tx)
	})
}