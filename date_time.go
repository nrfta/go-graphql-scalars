package scalars

import (
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/neighborly/go-errors"
)

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339, tmpStr)
	}
	return time.Time{}, errors.InvalidArgument.New("date should be ISO 8601 (2006-01-02T15:04:05Z07:00) formatted string")
}
