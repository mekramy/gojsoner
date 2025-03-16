package gojsoner

import (
	"slices"
	"strings"
)

func newFilter(filters ...string) *filter {
	return &filter{
		exclude: slices.DeleteFunc(
			filters,
			func(i string) bool {
				return strings.TrimSpace(i) == ""
			},
		),
	}

}

type filter struct {
	exclude []string
}

func (f *filter) shouldSkip(field string) bool {
	return field == "" || slices.Contains(f.exclude, field)
}
