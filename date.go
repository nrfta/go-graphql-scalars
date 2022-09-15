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
		result, err := time.Parse("2006-01-02", tmpStr)
		if err != nil {
			result2, err2 := time.Parse(time.RFC3339, tmpStr)

			if err2 != nil {
				return time.Time{}, err
			}

			return time.Parse("2006-01-02", result2.Format("2006-01-02"))
		}

		return result, nil
	}
	return time.Time{}, errors.InvalidArgument.New("date should be YYYY-MM-DD formatted string")
}
