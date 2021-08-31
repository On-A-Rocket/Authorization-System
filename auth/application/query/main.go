package query

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"gorm.io/gorm"
)

type Query struct {
	db     *gorm.DB
	config config.Interface
	// Login *LoginQueryHandler
	// Auth  *AuthQueryHandler
}

func NewQuery(db *gorm.DB, config config.Interface) *Query {
	return &Query{
		db:     db,
		config: config,
		// Login: newLoginQueryHandler(db, config),
		// Auth:  newAuthQueryHandler(config),
	}
}
