package pref

import "github.com/mattn/go-mastodon"

type Keys struct {
	Channel     string                `json:"channel"`
	App         *mastodon.Application `json:"app"`
	AccessToken string                `json:"access"`
}

func (instance *Keys) IsAuthenticated() bool {
	return instance.AccessToken != ""
}
