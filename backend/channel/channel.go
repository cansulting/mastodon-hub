package channel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"social/backend/constants"
	"strconv"
)

var errorMsg errorResponse

type errorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	Host string
}

func createHandler(host string) *Handler {
	return &Handler{
		Host: host,
	}
}

/// use to load timeline
/// @local true if only load the local timeline
func (instance *Handler) RetrieveLocalTimelines(local bool) (string, error) {
	url := getUrlFrom(instance.Host, constants.TIMELINES_ENDPOINT, "local="+strconv.FormatBool(local))
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("failed retrieving localtimelines. " + err.Error())
	}
	strRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed retrieving localtimelines. " + err.Error())
	}
	if err := json.Unmarshal(strRes, &errorMsg); err == nil {
		if errorMsg.Error != "" {
			return "", errors.New(errorMsg.Error)
		}
	}
	return string(strRes), nil
}
