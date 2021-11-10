package main

import (
	"social/backend/channel"
	"testing"
)

const testHost = "noc.social"

func TestLoad(t *testing.T) {
	handler := channel.Load(testHost)
	json, err := handler.RetrieveLocalTimelines()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json)
}
