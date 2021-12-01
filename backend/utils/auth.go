package utils

import (
	"context"
	"errors"
	"social/backend/pref"

	"github.com/mattn/go-mastodon"
)

const BASIC_SCOPE = "read write follow push" // basic scope

func Authenticate(channel string, username string, password string, prefDat *pref.Data) (*mastodon.Client, error) {
	keys := prefDat.GetChannelKeys(channel)
	// register app if not yet registered
	if keys == nil {
		_keys, err := RegisterApp(channel, prefDat)
		if err != nil {
			return nil, err
		}
		keys = _keys
	}
	config := &mastodon.Config{
		Server:       "https://" + channel,
		ClientID:     keys.App.ClientID,
		ClientSecret: keys.App.ClientSecret,
	}
	client := mastodon.NewClient(config)
	err := client.Authenticate(context.Background(), username, password)
	if err != nil {
		return nil, err
	}
	keys.AccessToken = config.AccessToken
	return client, nil
}

func SilentAuth(channel string, prefDat *pref.Data) (*mastodon.Client, error) {
	keys := prefDat.GetChannelKeys(channel)
	// register app if not yet registered
	if keys == nil {
		_keys, err := RegisterApp(channel, prefDat)
		if err != nil {
			return nil, err
		}
		keys = _keys
	}
	if keys.AccessToken == "" {
		return nil, errors.New("access token not found, authenticate with credential instead")
	}
	config := &mastodon.Config{
		Server:       "https://" + channel,
		ClientID:     keys.App.ClientID,
		ClientSecret: keys.App.ClientSecret,
		AccessToken:  keys.AccessToken,
	}
	client := mastodon.NewClient(config)
	return client, nil
}

func RegisterApp(channel string, prefDat *pref.Data) (*pref.Keys, error) {
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     "https://" + channel,
		Scopes:     BASIC_SCOPE,
		ClientName: "Elabox Mastodon",
	})
	if err != nil {
		return nil, err
	}
	keys := &pref.Keys{Channel: channel, App: app}
	prefDat.AddChannelKeys(keys)
	return keys, nil
}
