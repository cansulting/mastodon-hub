package mastodon

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"social/backend/utils"
	"strconv"
)

const reg_endpoint = "api/v1/apps" //
const access_token_endpoint = "oauth/token"
const redirect_uri = "urn:ietf:wg:oauth:2.0:oob" // only return the result instead of redirect
const BASIC_SCOPE = "read write follow push"     // basic scope

// use to register app and returns the app data
func RegApp(host string, clientName string, scope string) (*AppData, error) {
	url := utils.ConstructUrl(
		host,
		reg_endpoint,
		"scopes="+scope,
		"redirect_uris="+redirect_uri,
		"client_name="+clientName,
	)
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Status Code = " + strconv.Itoa(res.StatusCode) + ". " + res.Status)
	}
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := &AppData{}
	json.Unmarshal(resBytes, data)
	return data, nil
}

// retrieve access token based from app data
func RetrieveAccessToken(host string, scopes string, clientId string, clientSecret string) (*AccessData, error) {
	url := utils.ConstructUrl(
		host,
		access_token_endpoint,
		"scopes="+scopes,
		"redirect_uris="+redirect_uri,
		"client_id="+clientId,
		"client_secret="+clientSecret,
		"grant_type=client_credentials",
	)
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Status Code = " + strconv.Itoa(res.StatusCode) + ". " + res.Status)
	}
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := &AccessData{}
	json.Unmarshal(resBytes, data)
	return data, nil
}
