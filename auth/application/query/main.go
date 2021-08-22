package query

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"gorm.io/gorm"
)

type Query struct {
	Login *LoginQueryHandler
}

func NewQuery(db *gorm.DB, config config.Interface) *Query {
	return &Query{
		Login: newLoginQueryHandler(db, config),
	}
}
