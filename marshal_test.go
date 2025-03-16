package gojsoner_test

import (
	"testing"

	"github.com/mekramy/gojsoner"
)

func TestJsoner(t *testing.T) {
	type Book struct {
		Title string
		ISBN  string `json:"isbn"`
	}

	type Author struct {
		Name      string `json:"name"`
		Family    string `json:"family"`
		Age       int    `json:"age,string,omitempty"`
		IsMariage bool   `json:",string"`
		Books     []Book `json:"author_books"`
		Skills    []string
		Address   map[string]any
		// ignored fields
		PrivateField string `json:"-"`
		ignored      string
	}

	john := Author{
		Name:      "John",
		Family:    "Doe",
		Age:       0,
		IsMariage: false,
		Books: []Book{
			{Title: "Basics Of C", ISBN: "12345"},
			{Title: "Golang", ISBN: "88888"},
		},
		Skills: []string{"Web dev", "System programming", "IOT"},
		Address: map[string]any{
			"state": map[string]string{
				"country": "USA",
				"county":  "NY",
			},
			"city":   "NY city",
			"street": "ST. 23",
			"no":     13,
		},
		PrivateField: "Some private information",
		ignored:      "i'm ignored",
	}

	encoded, err := gojsoner.Marshal(
		john,
		"family",
		"Address.state.country",
		"author_books.isbn",
	)
	if err != nil {
		t.Error(err)
	}

	expected := `{"Address":{"city":"NY city","no":13,"state":{"county":"NY"},"street":"ST. 23"},"IsMariage":"false","Skills":["Web dev","System programming","IOT"],"author_books":[{"Title":"Basics Of C"},{"Title":"Golang"}],"name":"John"}`
	if string(encoded) != expected {
		t.Errorf("expected \n\t '%s'\n\t got '%s'", expected, string(encoded))
	}
}
