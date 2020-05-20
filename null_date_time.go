package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null/v8"
)

func MarshalNullDateTime(t null.Time) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return MarshalDateTime(t.Time)
}

func UnmarshalNullDateTime(v interface{}) (null.Time, error) {
	if v == nil {
		return null.TimeFromPtr(nil), nil
	}

	t, err := UnmarshalDateTime(v)
	return null.Time{Time: t}, err
}
