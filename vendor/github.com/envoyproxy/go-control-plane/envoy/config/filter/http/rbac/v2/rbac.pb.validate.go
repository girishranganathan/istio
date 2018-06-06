// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/rbac/v2/rbac.proto
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

// Validate checks the field values on RBAC with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *RBAC) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRules() == nil {
		return RBACValidationError{
			Field:  "Rules",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRules()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACValidationError{
				Field:  "Rules",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// RBACValidationError is the validation error returned by RBAC.Validate if the
// designated constraints aren't met.
type RBACValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RBACValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRBAC.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RBACValidationError{}

// Validate checks the field values on RBACPerRoute with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RBACPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Override.(type) {

	case *RBACPerRoute_Disabled:

		if m.GetDisabled() != true {
			return RBACPerRouteValidationError{
				Field:  "Disabled",
				Reason: "value must equal true",
			}
		}

	case *RBACPerRoute_Rbac:

		if m.GetRbac() == nil {
			return RBACPerRouteValidationError{
				Field:  "Rbac",
				Reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetRbac()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RBACPerRouteValidationError{
					Field:  "Rbac",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return RBACPerRouteValidationError{
			Field:  "Override",
			Reason: "value is required",
		}

	}

	return nil
}

// RBACPerRouteValidationError is the validation error returned by
// RBACPerRoute.Validate if the designated constraints aren't met.
type RBACPerRouteValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RBACPerRouteValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRBACPerRoute.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RBACPerRouteValidationError{}
