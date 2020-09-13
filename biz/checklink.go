package biz

import (
	"littlelink/model"
	"net/http"
)

func checkLink(link string) error {
	res, err := http.Head(link)
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		return nil
	}

	res, err = http.Get(link)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return model.ErrLinkFail
	}

	return nil
}
