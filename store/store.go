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


// GetUser calls the Db for a user lookup by coopID
func GetUser(ctx context.Context, ID string) (*User, error) {

	u, err := getUserByIdentifier(ctx, newGetUserByIDStatement(ID))
	if u != nil {
		return u, err
	}
	return nil, err
}

func getUserByIdentifier(ctx context.Context, stmt statement) (*User, error) {
	ctx, cancel, err := prepareForQuery(ctx)
	defer cancel()
	if err != nil {
		return nil, err
	}
	user := &User{}

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