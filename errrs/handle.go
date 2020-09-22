package errrs

import (
	"encoding/json"
	"littlelink/logger"
	"littlelink/model"
	"net/http"
)

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleWrongMethod(r *http.Request) (int, []byte) {
	go logger.ErrLog(model.LOW, model.ErrMethodNotAllowed.Error(), r)

	response := make(map[string]string)
	response["Error"] = model.ErrMethodNotAllowed.Error()
	response["Remediation"] = "Please use POST call for this API"
	response["POST Call Key"] = "link"

	resBytes, err := json.MarshalIndent(response, "", "	")
	Handle(err)

	return http.StatusMethodNotAllowed, resBytes
}

func HandleWrongQueryMethod(r *http.Request) (int, []byte) {
	response := make(map[string]string)
	response["Error"] = model.ErrQueryMethodNotAllowed.Error()
	response["Remediation"] = "Please use FORM BODY to send parameters this API"
	response["POST Call Key"] = "link"

	resBytes, err := json.MarshalIndent(response, "", "	")
	Handle(err)
	return http.StatusMethodNotAllowed, resBytes
}

func HandleLinkFail(r *http.Request, err string) (int, []byte) {
	go logger.ErrLog(model.MEDIUM, err, r)

	response := make(map[string]string)
	response["Error"] = model.ErrLinkFail.Error()
	response["Remediation"] = "Please verify the URL from your end once again"

	resBytes, e := json.MarshalIndent(response, "", "	")
	Handle(e)

	return http.StatusInternalServerError, resBytes
}

func HandleInternalError(r *http.Request, err string) (int, []byte) {
	go logger.ErrLog(model.HIGH, err, r)

	response := make(map[string]string)
	response["Error"] = model.ErrSomethingWentWrong.Error()

	resBytes, e := json.MarshalIndent(response, "", "	")
	Handle(e)

	return http.StatusInternalServerError, resBytes
}
