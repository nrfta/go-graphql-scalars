package scalars

import (
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/neighborly/go-errors"
)

func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format("2006-01-02")))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse("2006-01-02", tmpStr)
	}
	return time.Time{}, errors.InvalidArgument.New("date should be YYYY-MM-DD formatted string")
}
