# go-graphql-scalars ![](https://github.com/nrfta/go-graphql-scalars/workflows/CI/badge.svg)

A collection of GraphQL ([gqlgen](https://github.com/99designs/gqlgen/)) Scalars with [SQLBoiler's](https://github.com/volatiletech/sqlboiler) [null](https://github.com/volatiletech/null) support

## Install

```sh
go get -u "github.com/nrfta/go-graphql-scalars"
```

## Usage


```graphql
# scalars.graphql

"The `Date` is a date in the format YYYY-MM-DD"
scalar Date

"The `DateTime` is a date in the format ISO 8601 format: `2006-01-02T15:04:05Z07:00`"
scalar DateTime
```

```yml
# gqlgen.yml

schema:
  - "./*.graphql"

models:
  # -- Start Scalars --
  String:
    model:
      - github.com/99designs/gqlgen/graphql.String
      - github.com/nrfta/go-graphql-scalars.NullString
  Boolean:
    model:
      - github.com/99designs/gqlgen/graphql.Boolean
      - github.com/nrfta/go-graphql-scalars.NullBoolean
  Float:
    model:
      - github.com/99designs/gqlgen/graphql.Float
      - github.com/nrfta/go-graphql-scalars.NullFloat64
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int32
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/nrfta/go-graphql-scalars.NullInt
  Date:
    model:
      - github.com/nrfta/go-graphql-scalars.Date
      - github.com/nrfta/go-graphql-scalars.NullDate
  DateTime:
    model:
      - github.com/nrfta/go-graphql-scalars.DateTime
      - github.com/nrfta/go-graphql-scalars.NullDateTime
  # -- End Scalars --
```

## License

This project is licensed under the [MIT License](LICENSE.md).
