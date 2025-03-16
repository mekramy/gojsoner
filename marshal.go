package gojsoner

import "encoding/json"

// Marshal converts the input value (v) into JSON format with optional processing options.
// You must provide the "field.name" path for each field to be excluded.
func Marshal(v any, exclude ...string) ([]byte, error) {
	filter := newFilter(exclude...)
	filtered, err := mapper(v, "", filter)
	if err != nil {
		return nil, err
	}

	return json.Marshal(filtered)
}

// MarshalIndent converts the input value (v) into Indented JSON format with optional processing options.
// You must provide the "field.name" path for each field to be excluded.
func MarshalIndent(v any, indent string, exclude ...string) ([]byte, error) {
	filter := newFilter(exclude...)
	filtered, err := mapper(v, "", filter)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(filtered, "", indent)
}
