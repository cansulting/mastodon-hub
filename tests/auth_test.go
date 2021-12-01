package main

import (
	"social/backend/mastodon"
	"testing"
)

func TestRegApp(t *testing.T) {
	data, err := mastodon.RegApp("mastodon.social", "Test Elabox App", mastodon.BASIC_SCOPE)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestRetrieveAccessToken(t *testing.T) {
	data, err := mastodon.RetrieveAccessToken("mastodon.social", mastodon.BASIC_SCOPE, "HyrNcHT_p7pejMvVOSvvj2LvkotkBLlGL5Ef10-nrjQ", "ZHZZvAQ2D0o3sggnK7ET4GszTJFX3VD5ODtHbCl-Fec")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
