package oai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gitlab.glm.ai/ai-search/ai-e2e-chat/common/constant"
	video_types "gitlab.glm.ai/ai-search/ai-e2e-chat/video-sdk/types"
)

func newRandomID() string {
	u := uuid.New().String()
	return strings.ReplaceAll(u, "-", "")
}

func NewSessionID() string {
	return fmt.Sprintf("%s%s", "sess", newRandomID())
}

func NewEventID() string {
	return fmt.Sprintf("%s%s", "event", newRandomID())
}

func NewItemID() string {
	return fmt.Sprintf("%s%s", "item", newRandomID())
}

func NewResponseID() string {
	return fmt.Sprintf("%s%s", "resp", newRandomID())
}

func NewErrorEvent(code constant.RealtimeErrorCode, message string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    Error,
		Error: &EventError{
			Code:    code.String(),
			Message: message,
		},
	}
}

func NewSessionUpdatedEvent(session *Session, betaFields *BetaFields) Event {
	return Event{
		EventID: NewEventID(),
		Type:    SessionUpdated,
		Session: session,
	}
}

// 在这里初始化默认参数
func NewSessionCreatedEvent() Event {
	return Event{
		EventID: NewEventID(),
		Type:    SessionCreated,
		Session: &Session{
			ID:     NewSessionID(),
			Object: "realtime.session",
			Model:  "glm-realtime",
			Modalities: []Modality{
				ModalityAudio,
				ModalityText,
			},
			Instructions:      "",
			Voice:             "default",
			InputAudioFormat:  "wav",
			OutputAudioFormat: "pcm",
			// InputAudioTranscription: InputAudioTranscription{
			// 	Enabled: true,
			// 	Model:   "glm-realtime",
			// },
			Temperature: 0.05,
			BetaFields: &BetaFields{
				ChatMode: ChatModeAudio,
			},
		},
	}
}

func NewResponseCreatedEvent() Event {
	return Event{
		EventID: NewEventID(),
		Type:    ResponseCreated,
		Response: &Response{
			ID:     NewResponseID(),
			Object: ResponseObjectResponse,
			Status: ResponseStatusInProgress,
		},
	}
}

func NewInputBufferCommitedEvent(itemID string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    InputAudioBufferCommitted,
		ItemID:  itemID,
	}
}

func NewItemCreatedEvent(messageType ItemType, role ItemRole, status ItemStatus, content Content, itemID string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    ConversationItemCreated,
		Item: &Item{
			ID:     itemID,
			Object: ItemObjectRealTimeItem,
			Type:   messageType,
			Role:   role,
			Status: status,
			Content: []Content{
				content,
			},
		},
	}
}
func NewResponseDoneEvent(usage *Usage, status ResponseStatus, responseID string) Event {
	e := Event{
		EventID: NewEventID(),
		Type:    ResponseDone,
		Response: &Response{
			Status: status,
			ID:     responseID,
		},
	}
	if usage != nil {
		e.Response.Usage = *usage
	}
	return e
}

func NewResponseAudioDeltaEvent(delta string, extra *TTSExtra, responseID string) Event {
	e := Event{
		EventID:    NewEventID(),
		Type:       ResponseAudioDelta,
		Delta:      delta,
		ResponseID: responseID,
	}
	if extra != nil {
		e.Session = &Session{
			BetaFields: &BetaFields{
				TTSExtra: extra,
			},
		}
	}
	return e
}
func NewResponseTextDeltaEvent(delta string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    ResponseTextDelta,
		Delta:   delta,
	}
}

func NewInputAudioTranscriptionCompletedEvent(text string) Event {
	return Event{
		EventID:    NewEventID(),
		Type:       ConversationItemInputAudioTranscriptionCompleted,
		Transcript: text,
	}
}

func NewResponseAudioTranscriptDeltaEvent(delta string, isLast bool, responseID string) Event {
	return Event{
		EventID:    NewEventID(),
		Type:       ResponseAudioTranscriptDelta,
		Delta:      delta,
		ResponseID: responseID,
	}
}

func NewInputAudioBufferSpeechStartedEvent() Event {
	return Event{
		EventID: NewEventID(),
		Type:    InputAudioBufferSpeechStarted,
	}
}

func NewInputAudioBufferSpeechStoppedEvent() Event {
	return Event{
		EventID: NewEventID(),
		Type:    InputAudioBufferSpeechStopped,
	}
}

func NewResponseFunctionCallArgumentsDoneEvent(name string, arguments string, responseID string) Event {
	return Event{
		EventID:    NewEventID(),
		Type:       ResponseFunctionCallArgumentsDone,
		Name:       name,
		Arguments:  arguments,
		ResponseID: responseID,
	}
}

// func NewUsageEvent(usage Usage) Event {
// 	return Event{
// 		EventID: NewEventID(),
// 		Type:    ResponseCreate,
// 		Response: &Response{
// 			Usage: usage,
// 		},
// 	}
// }

// 定制化事件 OpenAI没有
func NewQinYanResponseFunctionCallEvent(name string, description string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    ResponseFunctionCallSimpleBrowserEvent,
		Name:    name,
		Session: &Session{
			BetaFields: &BetaFields{
				SimpleBrowser: &SimpleBrowser{
					Description: description,
				},
			},
		},
	}
}

// 当前不会使用
func NewQinYanResponseFunctionCallOutputEvent(name string, chunk video_types.ResultChunk) Event {
	jsonStr, err := json.Marshal(chunk.SearchMeta)
	if err != nil {
		return NewErrorEvent("json_marshal_error", "NewQinYanResponseFunctionCallOutputEvent:"+err.Error())
	}
	return Event{
		EventID: NewEventID(),
		Type:    ResponseFunctionCallSimpleBrowserResultEvent,
		Name:    name,
		Session: &Session{
			BetaFields: &BetaFields{
				SimpleBrowser: &SimpleBrowser{
					SearchMeta:   string(jsonStr),
					Meta:         chunk.Meta,
					TextCitation: chunk.Text,
				},
			},
		},
	}
}

func NewConversationItemMessageEvent(content Content, role ItemRole) Event {
	return Event{
		EventID: NewEventID(),
		Type:    ConversationItemCreate,
		Item: &Item{
			Type: ItemTypeMessage,
			Role: role,
			Content: []Content{
				content,
			},
		},
	}
}

func NewConversationItemToolOutputEvent(content Content, role ItemRole, output string) Event {
	return Event{
		EventID: NewEventID(),
		Type:    ConversationItemCreate,
		Item: &Item{
			Type:    ItemTypeFunctionCallOutput,
			Role:    role,
			Content: []Content{content},
			Output:  &output,
		},
	}
}
