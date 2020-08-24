package dblayer

import (
	"littlelink/model"

	"github.com/xujiajun/nutsdb"
)

func InitDB() (*nutsdb.DB, error) {
	opt := nutsdb.DefaultOptions
	opt.Dir = model.DBSTRING
	return nutsdb.Open(opt)
}
