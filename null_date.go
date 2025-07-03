package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aarondl/null/v8"
)

func MarshalNullDate(t null.Time) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return MarshalDate(t.Time)
}

func UnmarshalNullDate(v interface{}) (null.Time, error) {
	if v == nil {
		return null.TimeFromPtr(nil), nil
	}

	t, err := UnmarshalDate(v)
	return null.Time{Time: t}, err
}
