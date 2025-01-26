package oai

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChatMode(t *testing.T) {
	session := &Session{
		BetaFields: &BetaFields{},
	}
	chatMode := GetChatMode(context.Background(), session)
	assert.Equal(t, chatMode, ChatModeAudio)

	session = &Session{
		BetaFields: nil,
	}
	chatMode = GetChatMode(context.Background(), session)
	assert.Equal(t, chatMode, ChatModeAudio)
}
