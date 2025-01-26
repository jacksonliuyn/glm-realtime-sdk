package oai

import (
	"context"
	"encoding/json"

	"gitlab.glm.ai/ai-search/ai-e2e-chat/common/constant"
	"gitlab.glm.ai/ai-search/ai-e2e-chat/common/logs"
)

type Event struct {
	EventID         string      `json:"event_id,omitempty"`
	Type            EventType   `json:"type"`
	Session         *Session    `json:"session,omitempty"`
	Audio           string      `json:"audio,omitempty"`
	Response        *Response   `json:"response,omitempty"`
	ItemID          string      `json:"item_id,omitempty"`
	PreviousItemID  string      `json:"previous_item_id,omitempty"`
	ResponseID      string      `json:"response_id,omitempty"`
	OutputIndex     int         `json:"output_index,omitempty"`
	ContentIndex    int         `json:"content_index,omitempty"`
	Delta           string      `json:"delta"`
	Item            *Item       `json:"item,omitempty"`
	ClientTimestamp int64       `json:"client_timestamp,omitempty"`
	Transcript      string      `json:"transcript,omitempty"`
	Name            string      `json:"name,omitempty"`
	Arguments       string      `json:"arguments,omitempty"`
	VideoFrame      []byte      `json:"video_frame,omitempty"`
	Instructions    string      `json:"instructions,omitempty"`
	Error           *EventError `json:"error,omitempty"`
	// BetaFields      *BetaFields `json:"beta_fields,omitempty"`
}

type EventError struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Event) GetConversationID() string {
	if e == nil {
		return ""
	}

	if e.Session == nil {
		return ""
	}
	if e.Session.ID == "" {
		return ""
	}
	return e.Session.ID
}

func (e *Event) GetMessageID() string {
	if e == nil {
		return ""
	}
	if e.ItemID != "" {
		return e.ItemID
	}
	if e.Item == nil {
		return ""
	}
	if e.Item.ID == "" {
		return ""
	}
	return e.Item.ID
}

func (e *Event) ToJson() string {
	json, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(json)
}

type Modality string
type ChatMode string

const (
	ModalityText  Modality = "text"
	ModalityAudio Modality = "audio"
	ModalityVideo Modality = "video"
)

const (
	ChatModeAudio          ChatMode = "audio"
	ChatModeVideoPassive   ChatMode = "video_passive"
	ChatModeVideoProactive ChatMode = "video_preactive"
)

type Session struct {
	ID                      string                  `json:"id"`
	Object                  string                  `json:"object"`
	Model                   string                  `json:"model"`
	Modalities              []Modality              `json:"modalities"`
	Instructions            string                  `json:"instructions"`
	Voice                   string                  `json:"voice"`
	InputAudioFormat        string                  `json:"input_audio_format"`
	OutputAudioFormat       string                  `json:"output_audio_format"`
	InputAudioTranscription *InputAudioTranscription `json:"input_audio_transcription"`
	TurnDetection           *TurnDetection          `json:"turn_detection"`
	Tools                   []Tool                  `json:"tools"`
	ToolChoice              string                  `json:"tool_choice"`
	Temperature             float64                 `json:"temperature"`
	MaxOutputTokens         any                     `json:"max_output_tokens"` // "inf" or int
	BetaFields              *BetaFields             `json:"beta_fields"`
}

type InputAudioTranscription struct {
	Enabled bool   `json:"enabled"`
	Model   string `json:"model"`
}

type TurnDetection struct {
	Type              string  `json:"type"`
	Threshold         float64 `json:"threshold"`
	PrefixPaddingMs   int     `json:"prefix_padding_ms"`
	SilenceDurationMs int     `json:"silence_duration_ms"`
}

type TTSExtra struct {
	Index    int    `json:"index"`
	SubIndex int    `json:"sub_index"`
	SubText  string `json:"sub_text"`
	IsEnd    bool   `json:"is_end"`
}

type SimpleBrowser struct {
	Description  string `json:"description"`
	SearchMeta   string `json:"search_meta"`
	Meta         string `json:"meta"`
	TextCitation string `json:"text_citation"`
}

type BetaFields struct {
	ChatMode        ChatMode       `json:"chat_mode,omitempty"`
	ImageSizeX      int            `json:"image_size_x,omitempty"`
	ImageSizeY      int            `json:"image_size_y,omitempty"`
	FPS             int            `json:"fps,omitempty"`
	TTSSource       string         `json:"tts_source,omitempty"`
	TTSExtra        *TTSExtra      `json:"tts_extra,omitempty"`
	InputSampleRate int            `json:"input_sample_rate,omitempty"` // 输入采样率
	SimpleBrowser   *SimpleBrowser `json:"simple_browser,omitempty"`
	IsLastText      bool           `json:"is_last_text"`
	MessageID       string         `json:"message_id,omitempty"`  // 清言传过来的消息ID
	AutoSearch      *bool          `json:"auto_search,omitempty"` // 是否自动搜索
}

func GetChatMode(ctx context.Context, session *Session) ChatMode {
	if session.BetaFields != nil && session.BetaFields.ChatMode != "" {
		return session.BetaFields.ChatMode
	}
	logs.FormatInfo(ctx, "GetChatMode 没有设置 chat_mode, 使用默认值: %s", ChatModeAudio)
	return ChatModeAudio
}

func GetTTSSource(session *Session) string {
	if session.BetaFields != nil && (session.BetaFields.TTSSource == constant.TTS_SOURCE_ZHIPU ||
		session.BetaFields.TTSSource == constant.TTS_SOURCE_HUOSHAN || session.BetaFields.TTSSource == constant.TTS_SOURCE_E2E) {
		return session.BetaFields.TTSSource
	}
	return ""
}

func IsAutoSearch(session *Session) bool {
	//  auto search 不传 默认开启, 传了就读取
	if session.BetaFields != nil && session.BetaFields.AutoSearch != nil {
		return *session.BetaFields.AutoSearch
	}
	return true
}
