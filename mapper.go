package gojsoner

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func mapper(v any, path string, filters *filter) (any, error) {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return nil, nil
	}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, nil
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		result := make([]any, 0)
		for i := 0; i < value.Len(); i++ {
			res, err := mapper(value.Index(i).Interface(), path, filters)
			if err != nil {
				return nil, err
			}
			result = append(result, res)
		}
		return result, nil
	case reflect.Map:
		result := make(map[string]any)
		for _, k := range value.MapKeys() {
			key := fmt.Sprint(k.Interface())
			if filters.shouldSkip(pathJoiner(path, key)) {
				continue
			}

			res, err := mapper(
				value.MapIndex(k).Interface(),
				pathJoiner(path, key),
				filters,
			)
			if err != nil {
				return nil, err
			}

			result[key] = res
		}
		return result, nil
	case reflect.Struct:
		encoded, err := json.Marshal(value.Interface())
		if err != nil {
			return nil, err
		}

		var decoded map[string]any
		err = json.Unmarshal(encoded, &decoded)
		if err != nil {
			return nil, err
		}

		return mapper(decoded, path, filters)
	}

	return v, nil
}

func pathJoiner(root, name string) string {
	if root == "" {
		return name
	} else {
		return root + "." + name
	}
}
