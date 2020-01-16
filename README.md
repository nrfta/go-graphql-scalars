# go-graphql-scalars

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

```

## License

This project is licensed under the [MIT License](LICENSE.md).
