package scalars

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
)

// JSON represents a custom scalar for raw JSON data
type JSON json.RawMessage

// MarshalJSON implements the graphql.Marshaler interface for JSON
func MarshalJSON(j JSON) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if _, err := w.Write(j); err != nil {
			slog.Error("JSON Marshal scalar: fail while writing data", "error", err)
		}
	})
}

// UnmarshalJSON implements the graphql.Unmarshaler interface for JSON
func UnmarshalJSON(v interface{}) (JSON, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		// Marshal the map into raw JSON bytes
		data, err := json.Marshal(val)
		if err != nil {
			return nil, fmt.Errorf("JSON Unmarshal scalar: fail while marshaling map to json: %w", err)
		}
		return JSON(data), nil
	case string:
		// Assume the string is already JSON
		return JSON(val), nil
	default:
		return nil, fmt.Errorf("JSON Unmarshal scalar: unexpected type: %T", v)
	}
}
