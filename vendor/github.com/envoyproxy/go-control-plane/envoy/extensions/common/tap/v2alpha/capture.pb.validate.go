// Code generated by protoc-gen-validate
// source: envoy/extensions/common/tap/v2alpha/capture.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on Connection with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Connection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetLocalAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				Field:  "LocalAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				Field:  "RemoteAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ConnectionValidationError is the validation error returned by
// Connection.Validate if the designated constraints aren't met.
type ConnectionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ConnectionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnection.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ConnectionValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				Field:  "Timestamp",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.EventSelector.(type) {

	case *Event_Read_:

		if v, ok := interface{}(m.GetRead()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					Field:  "Read",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Event_Write_:

		if v, ok := interface{}(m.GetWrite()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					Field:  "Write",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = EventValidationError{}

// Validate checks the field values on Trace with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trace) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TraceValidationError{
				Field:  "Connection",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceValidationError{
					Field:  fmt.Sprintf("Events[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// TraceValidationError is the validation error returned by Trace.Validate if
// the designated constraints aren't met.
type TraceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TraceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrace.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TraceValidationError{}

// Validate checks the field values on Event_Read with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event_Read) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Data

	return nil
}

// Event_ReadValidationError is the validation error returned by
// Event_Read.Validate if the designated constraints aren't met.
type Event_ReadValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Event_ReadValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent_Read.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Event_ReadValidationError{}

// Validate checks the field values on Event_Write with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Event_Write) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Data

	// no validation rules for EndStream

	return nil
}

// Event_WriteValidationError is the validation error returned by
// Event_Write.Validate if the designated constraints aren't met.
type Event_WriteValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Event_WriteValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent_Write.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Event_WriteValidationError{}
