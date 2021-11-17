<h1>MSSQL DB example with dd-trace/sql</h1>

Repo for demonstrating how using
[datadog dd-trace/sql tracing](https://pkg.go.dev/gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql) with a 
[mssql-driver](https://github.com/denisenkom/go-mssqldb) returns error
for Named Parameters.

<h3>Running the database</h3>

Create a database image by running `./build_image.sh` in `./docker/`. 
Then run the database by doing `docker run -it -p 1433:1433 example/mssql-db:datadog`.


