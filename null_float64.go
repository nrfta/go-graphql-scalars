package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aarondl/null/v8"
)

func MarshalNullFloat64(ns null.Float64) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalFloat(ns.Float64)
}

func UnmarshalNullFloat64(v interface{}) (null.Float64, error) {
	if v == nil {
		return null.Float64{Valid: false}, nil
	}
	s, err := graphql.UnmarshalFloat(v)
	return null.Float64{Float64: s}, err
}
