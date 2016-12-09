package test

import (
	"testing"
	"shoutshoutshout/libraries"
)

var watson = libraries.NewWatson()

func TestWatson(t *testing.T) {
	data := watson.ConversationApi("疲れた")

	if len(data) < 1 {
		t.Error("api connection failed.")
	}

	if data != "がんばれ（´・ω・｀）" {
		t.Error("api returned unmatch word.")
	}
}
