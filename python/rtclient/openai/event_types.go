package oai

type EventType string

const (
	// Session events
	SessionUpdate  EventType = "session.update"
	SessionCreated EventType = "session.created"
	SessionUpdated EventType = "session.updated"

	// Input audio buffer events
	InputAudioBufferAppend        EventType = "input_audio_buffer.append"
	InputAudioBufferCommit        EventType = "input_audio_buffer.commit"
	InputAudioBufferPreCommit     EventType = "input_audio_buffer.pre_commit"
	InputAudioBufferClear         EventType = "input_audio_buffer.clear"
	InputAudioBufferCommitted     EventType = "input_audio_buffer.committed"
	InputAudioBufferCleared       EventType = "input_audio_buffer.cleared"
	InputAudioBufferSpeechStarted EventType = "input_audio_buffer.speech_started"
	InputAudioBufferSpeechStopped EventType = "input_audio_buffer.speech_stopped"

	// video events
	InputVideoFrameAppend EventType = "input_audio_buffer.append_video_frame"

	// Conversation item events
	ConversationItemCreate                           EventType = "conversation.item.create"
	ConversationItemTruncate                         EventType = "conversation.item.truncate"
	ConversationItemDelete                           EventType = "conversation.item.delete"
	ConversationItemCreated                          EventType = "conversation.item.created"
	ConversationItemInputAudioTranscriptionCompleted EventType = "conversation.item.input_audio_transcription.completed"
	ConversationItemInputAudioTranscriptionFailed    EventType = "conversation.item.input_audio_transcription.failed"
	ConversationItemTruncated                        EventType = "conversation.item.truncated"
	ConversationItemDeleted                          EventType = "conversation.item.deleted"

	// Response events
	ResponseCreate                     EventType = "response.create"
	ResponseCancel                     EventType = "response.cancel"
	ResponseCreated                    EventType = "response.created"
	ResponseDone                       EventType = "response.done"
	ResponseOutputItemAdded            EventType = "response.output_item.added"
	ResponseOutputItemDone             EventType = "response.output_item.done"
	ResponseContentPartAdded           EventType = "response.content_part.added"
	ResponseContentPartDone            EventType = "response.content_part.done"
	ResponseTextDelta                  EventType = "response.text.delta"
	ResponseTextDone                   EventType = "response.text.done"
	ResponseAudioTranscriptDelta       EventType = "response.audio_transcript.delta"
	ResponseAudioTranscriptDone        EventType = "response.audio_transcript.done"
	ResponseAudioDelta                 EventType = "response.audio.delta"
	ResponseAudioDone                  EventType = "response.audio.done"
	ResponseFunctionCallArgumentsDelta EventType = "response.function_call_arguments.delta"
	ResponseFunctionCallArgumentsDone  EventType = "response.function_call_arguments.done"
	// 定制化事件
	ResponseFunctionCallSimpleBrowserEvent EventType = "response.function_call.simple_browser"
	ResponseFunctionCallSimpleBrowserResultEvent EventType = "response.function_call.simple_browser.result"
	// Server events
	ServerEvents EventType = "Server events"
	Error        EventType = "error"

	// Rate limits
	RateLimitsUpdated EventType = "rate_limits.updated"
)
