package router

import (
	"littlelink/biz"
	"littlelink/logger"
	"net/http"
	"time"
)

func (br *BaseRouter) SetSmallRoute(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	biz.SetHeaders(w)
	status, res := biz.SetSmall(r, br.db)
	w.WriteHeader(status)
	w.Write(res)

	go logger.CallLog(r, time.Since(start))
}
