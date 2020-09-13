package dblayer

import (
	"littlelink/model"
	"math"

	"github.com/xujiajun/nutsdb"
)

func SetDB(key, value string, db *nutsdb.DB) error {
	err := db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(model.BUCKETSTRING, []byte(key), []byte(value), math.MaxUint32)
	})
	return err
}
