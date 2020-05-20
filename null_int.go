package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null/v8"
)

func MarshalNullInt(ns null.Int) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalInt(ns.Int)
}

func UnmarshalNullInt(v interface{}) (null.Int, error) {
	if v == nil {
		return null.Int{Valid: false}, nil
	}
	s, err := graphql.UnmarshalInt(v)
	return null.Int{Int: s}, err
}
