package main

import (
	"context"
	"database/sql"
	"fmt"
	mssql "github.com/denisenkom/go-mssqldb"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	"log"
	"mssql-datadog-sqltrace-example/store"

	//Uncomment this to try with original "database/sql"
	//_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	ctx := context.Background()
	var err error

	//Comment out the line below and switch to "sql.Open(...)" to try with original "database/sql"
	sqltrace.Register("sqlserver", &mssql.Driver{})
	store.Database.Db, err = sql.Open("sqlserver", "sqlserver://sa:myPassw0rd@localhost:1433?database=ddExample")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := store.Database.Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	usr, err := store.GetUser(ctx, "1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User found: %v", usr)
}