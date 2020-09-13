package biz

import (
	"fmt"
	"littlelink/dblayer"
	"littlelink/errrs"
	"net/http"

	"github.com/xujiajun/nutsdb"
	"gopkg.in/square/go-jose.v2/json"
)

func SetSmall(r *http.Request, db *nutsdb.DB) (int, []byte) {
	if r.Method != http.MethodPost {
		return errrs.HandleWrongMethod(r)
	}

	link := r.FormValue("link")
	if err := checkLink(link); err != nil {
		return errrs.HandleLinkFail(r, err.Error())
	}
	sid := linkHash(link)

	value, _ := dblayer.GetDB(sid, db)
	if value == "" {
		if err := dblayer.SetDB(sid, link, db); err != nil {
			return errrs.HandleInternalError(r, err.Error())
		}
	}

	response := make(map[string]string)
	response["short"] = fmt.Sprintf("%s", sid)
	resBytes, err := json.Marshal(response)
	errrs.Handle(err)

	return http.StatusOK, resBytes
}
