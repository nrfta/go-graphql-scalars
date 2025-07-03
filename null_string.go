package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aarondl/null/v8"
)

func MarshalNullString(ns null.String) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

func UnmarshalNullString(v interface{}) (null.String, error) {
	if v == nil {
		return null.String{Valid: false}, nil
	}
	s, err := graphql.UnmarshalString(v)
	return null.String{String: s}, err
}
