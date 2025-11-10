package graphqlclient

import (
	"encoding/json"
	"errors"
	"strings"
)

type (
	// https://spec.graphql.org/draft/#sec-Errors
	Error []ErrorEntry

	ErrorEntry struct {
		Message    string                     `json:"message"`
		Path       []string                   `json:"path"`
		Locations  []Location                 `json:"locations"`
		Extensions map[string]json.RawMessage `json:"extensions"`
	}

	Location struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	}
)

var (
	ErrExtensionNotFound = errors.New("extension not found")
)

func (err Error) Unwrap() []error {
	var errs []error
	for _, entry := range err {
		errs = append(errs, entry)
	}
	return errs
}

func (err Error) Error() string {
	var b strings.Builder
	for i, entry := range err {
		_, _ = b.WriteString(entry.Error())

		if i < len(err)-1 {
			_ = b.WriteByte('\n')
		}
	}
	return b.String()
}

func (err Error) ExtensionBool(key string) (bool, bool) {
	for _, entry := range err {
		if val, ok := entry.ExtensionBool(key); ok {
			return val, ok
		}
	}
	return false, false
}

func (err Error) ExtensionFloat64(key string) (float64, bool) {
	for _, entry := range err {
		if val, ok := entry.ExtensionFloat64(key); ok {
			return val, ok
		}
	}
	return 0.0, false
}

func (err Error) ExtensionInt(key string) (int, bool) {
	for _, entry := range err {
		if val, ok := entry.ExtensionInt(key); ok {
			return val, ok
		}
	}
	return 0, false
}

func (err Error) ExtensionString(key string) (string, bool) {
	for _, entry := range err {
		if val, ok := entry.ExtensionString(key); ok {
			return val, ok
		}
	}
	return "", false
}

func (err Error) ExtensionEqualBool(key string, value bool) bool {
	for _, entry := range err {
		if val, ok := entry.ExtensionBool(key); ok && val == value {
			return true
		}
	}
	return false
}

func (err Error) ExtensionEqualInt(key string, value int) bool {
	for _, entry := range err {
		if val, ok := entry.ExtensionInt(key); ok && val == value {
			return true
		}
	}
	return false
}

func (err Error) ExtensionEqualString(key, value string) bool {
	for _, entry := range err {
		if val, ok := entry.ExtensionString(key); ok && val == value {
			return true
		}
	}
	return false
}

func (err Error) ExtensionUnmarshal(key string, in any) error {
	for _, entry := range err {
		if err := entry.ExtensionUnmarshal(key, in); err != ErrExtensionNotFound {
			return err
		}
	}
	return ErrExtensionNotFound
}

func (entry ErrorEntry) Error() string {
	return entry.Message
}

func (entry ErrorEntry) ExtensionBool(key string) (bool, bool) {
	var val bool
	if err := entry.ExtensionUnmarshal(key, &val); err != nil {
		return val, false
	}
	return val, true
}

func (entry ErrorEntry) ExtensionFloat64(key string) (float64, bool) {
	var val float64
	if err := entry.ExtensionUnmarshal(key, &val); err != nil {
		return 0, false
	}
	return val, true
}

func (entry ErrorEntry) ExtensionInt(key string) (int, bool) {
	var val int
	if err := entry.ExtensionUnmarshal(key, &val); err != nil {
		return 0, false
	}
	return val, true
}

func (entry ErrorEntry) ExtensionString(key string) (string, bool) {
	var val string
	if err := entry.ExtensionUnmarshal(key, &val); err != nil {
		return "", false
	}
	return val, true
}

func (entry ErrorEntry) ExtensionUnmarshal(key string, in any) error {
	if val, ok := entry.Extensions[key]; ok {
		return json.Unmarshal(val, in)
	}
	return ErrExtensionNotFound
}
