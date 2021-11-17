package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var Database struct {
	Db *sql.DB
}

type User struct {
	ID     string
	Name  string
	Email string
	Phone  string
}

func GetUser(ctx context.Context, ID string) (*User, error) {
	ctx, cancel, err := prepareForQuery(ctx)
	defer cancel()
	if err != nil {
		return nil, err
	}
	user := &User{}

	stmt := newGetUserByIDStatement(ID)

	err = Database.Db.QueryRowContext(ctx, stmt.getQuery(), stmt.getNamed()...).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func prepareForQuery(ctx context.Context) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	if Database.Db == nil {
		return ctx, cancel, fmt.Errorf("no DB connection found")
	}
	return ctx, cancel, nil
}