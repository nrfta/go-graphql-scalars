package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null"
)

func MarshalNullBoolean(ns null.Bool) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalBoolean(ns.Bool)
}

func UnmarshalNullBoolean(v interface{}) (null.Bool, error) {
	if v == nil {
		return null.Bool{Valid: false}, nil
	}
	s, err := graphql.UnmarshalBoolean(v)
	return null.Bool{Bool: s}, err
}