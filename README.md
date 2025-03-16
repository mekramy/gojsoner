# API Documentation

The GoJSONer library provides utility functions to convert data into JSON format with optional field exclusions. It allows you to specify which fields should be excluded from the output using a dot notation path (`"field", "nested.field", ...`). The main functions provided are Marshal and MarshalIndent, which can generate standard and indented JSON output, respectively.

## Marshal

Converts the input value `v` into JSON format. The exclude parameter allows specifying the paths of fields to be excluded from the final output.

```go
func Marshal(v any, exclude ...string) ([]byte, error)
```

```go
package main

import (
    "fmt"
    "github.com/mekramy/gojsoner"
)

type Author struct {
    Name string `json:"name"`
    Family string `json:"family"`
    Books []string `json:"books"`
}

func main() {
    author := Author{
        Name: "John Doe",
        Family: "Doe",
        Books: []string{"Go Programming", "Learning JSON"},
    }

    encoded, err := gojsoner.Marshal(author, "family")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(encoded))
    //  {
    //      "name": "John Doe",
    //      "books": ["Go Programming", "Learning JSON"]
    //  }
}
```

## MarshalIndent

Converts the input value `v` into an indented JSON format. Like Marshal, the exclude parameter allows specifying field paths to exclude from the resulting JSON.

```go
func MarshalIndent(v any, indent string, exclude ...string) ([]byte, error)
```

```go
package main

import (
    "fmt"
    "github.com/mekramy/gojsoner"
)

type Author struct {
    Name string `json:"name"`
    Family string `json:"family"`
    Books []string `json:"books"`
}

func main() {
    author := Author{
        Name: "Jane Doe",
        Family: "Doe",
        Books: []string{"Go Programming", "Learning JSON"},
    }

    encoded, err := gojsoner.MarshalIndent(author, "  ", "family")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(encoded))
    //  {
    //      "name": "Jane Doe",
    //      "books": [
    //          "Go Programming",
    //          "Learning JSON"
    //      ]
    //  }

}

```

### Conclusion

The GoJSONer library allows for flexible and efficient manipulation of Go data structures into JSON, with the ability to exclude specific fields using dot notation paths. This is particularly useful for sanitizing or excluding sensitive data, or simply tailoring the output for different use cases.

For more advanced customization, the library also supports nested field exclusions and structured filtering of complex objects.
