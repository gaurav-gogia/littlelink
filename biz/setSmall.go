package biz

import (
	"littlelink/dblayer"
	"littlelink/errrs"
	"net/http"

	"encoding/json"

	"github.com/xujiajun/nutsdb"
)

func SetSmall(r *http.Request, db *nutsdb.DB) (int, []byte) {
	if r.Method != http.MethodPost {
		return errrs.HandleWrongMethod(r)
	}

	link := r.PostFormValue("link")
	if link == "" {
		return errrs.HandleWrongQueryMethod(r)
	}

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
	response["short"] = sid
	resBytes, err := json.Marshal(response)
	errrs.Handle(err)

	return http.StatusOK, resBytes
}
