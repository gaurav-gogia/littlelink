package dblayer

import (
	"littlelink/model"

	"github.com/xujiajun/nutsdb"
)

func GetDB(key string, db *nutsdb.DB) (string, error) {
	var value string

	err := db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(model.BUCKETSTRING, []byte(key))
		if err != nil {
			return err
		}

		value = string(entry.Value)
		return nil
	})

	return value, err
}
