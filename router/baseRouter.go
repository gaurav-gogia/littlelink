package router

import (
	"github.com/xujiajun/nutsdb"
)

// BaseRouter structure defines the routers that helps all routers use linkdb work well with database
type BaseRouter struct {
	db *nutsdb.DB
}

// NewBaseRouter method returns a new base router
func NewBaseRouter(db *nutsdb.DB) *BaseRouter {
	return &BaseRouter{db: db}
}
