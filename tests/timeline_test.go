package main

import (
	"social/backend/channel"
	"testing"
)

const testHost = "noc.social"

// test loading timeline content
func TestLoad(t *testing.T) {
	handler := channel.Load(testHost)
	json, err := handler.RetrieveLocalTimelines(true)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json)
}
