package sourcetransformer

import (
	"fmt"
	"time"
)

var (
	DROP = fmt.Sprintf("%U__DROP__", '\\') // U+005C__DROP__
	// Watermark are at millisecond granularity, hence we use epoch(0) - 1 to indicate watermark is not available.
	// eventTimeForDrop is used to indicate that the message is dropped hence, excluded from watermark calculation
	eventTimeForDrop = time.Unix(0, -int64(time.Millisecond))
)

// Message is used to wrap the data return by SourceTransformer functions.
// Compared with Message of other UDFs, source transformer Message contains one more field,
// the event time, usually extracted from the payload.
type Message struct {
	value     []byte
	eventTime time.Time
	keys      []string
	tags      []string
}

// NewMessage creates a Message with eventTime and value
func NewMessage(value []byte, eventTime time.Time) Message {
	return Message{value: value, eventTime: eventTime}
}

// WithKeys is used to assign the keys to message
func (m Message) WithKeys(keys []string) Message {
	m.keys = keys
	return m
}

// WithTags is used to assign the tags to message
// tags will be used for conditional forwarding
func (m Message) WithTags(tags []string) Message {
	m.tags = tags
	return m
}

// EventTime returns message eventTime
func (m Message) EventTime() time.Time {
	return m.eventTime
}

// Keys returns message keys
func (m Message) Keys() []string {
	return m.keys
}

// Value returns message value
func (m Message) Value() []byte {
	return m.value
}

// Tags returns message tags
func (m Message) Tags() []string {
	return m.tags
}

// MessageToDrop creates a Message to be dropped
func MessageToDrop() Message {
	return Message{eventTime: eventTimeForDrop, value: []byte{}, tags: []string{DROP}}
}

type Messages []Message

// MessagesBuilder returns an empty instance of Messages
func MessagesBuilder() Messages {
	return Messages{}
}

// Append appends a Message
func (m Messages) Append(msg Message) Messages {
	m = append(m, msg)
	return m
}

// Items returns the Message list
func (m Messages) Items() []Message {
	return m
}
