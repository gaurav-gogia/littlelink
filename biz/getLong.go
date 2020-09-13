package biz

import (
	"net/http"

	"github.com/xujiajun/nutsdb"
)

func GetLong(r *http.Request, db *nutsdb.DB) (int, []byte) {
	var link string
	// ToDo: complete this function
	return http.StatusOK, []byte(link)
}
