package store

import (
	"database/sql"
)

// statement interface to make sql query with named parameters
type statement interface {
	getQuery() string
	getNamed() []interface{}
}

// selectUserByProfileKey should conform to statement interface
type namedQuery struct {
	query     string
	namedArgs []sql.NamedArg
}

func newGetUserByIDStatement(ID string) *namedQuery {

	q := `
		SELECT u.Id, u.Name, u.Email, u.PhoneNumber
		FROM Users AS u 
		WHERE u.Id=@Id
		`
	return &namedQuery{
		query:     q,
		namedArgs: []sql.NamedArg{sql.Named("Id", ID)},
	}
}


func (s *namedQuery) getQuery() string {
	return s.query
}

func (s *namedQuery) getNamed() []interface{} {
	var namedParams []interface{}
	for _, p := range s.namedArgs {
		namedParams = append(namedParams, p)
	}
	return namedParams
}
