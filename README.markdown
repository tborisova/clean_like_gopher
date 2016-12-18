# Clean like gopher

Clean like gopher is inspired from [database_cleaner](https://github.com/DatabaseCleaner/database_cleaner). The purpose of clean
like gopher is to ensure that the DB is clean between tests.

[![GoDoc](https://godoc.org/github.com/tborisova/clean_like_gopher?status.png)](https://godoc.org/github.com/tborisova/clean_like_gopher)

## TODO:
 * Add transaction for mysql
 * Implement redis cleaner
 * Use ginkgo for test
 * Write better tests for different strategies
 
 
## Supported drivers
   * [mysql](github.com/go-sql-driver/mysql)
   * [postgres](github.com/lib/pq)
   * [redis](https://github.com/go-redis/redis) - in future development
   * [mongo](https://labix.org/mgo)

## Install:
```
  go get "github.com/tborisova/clean_like_gopher"
```
## Usage

```go
  import 'github/tborisova/clean_like_gopher'
  ...
  
  options := map[string]string{"host": "localhost", "dbName": "test", "port": "27017"}
  m := clean_like_gopher.NewCleaningGopher("mongo", options) // clean collection 'test' using mongo driver and truncation strategy
  ...
  dirty db
  ...
  options = map[string]string{"strategy": "truncation"}
  m.Clean(options) // clean all tables with truncation
  m.Close() // after all specs or after each spec
```

## Examples: 
 
  * [clean-like-gopher-example](https://github.com/tborisova/examples-cleaning-gopher)
  
Availabe strategies:

  * For mysql/postgres:
    * truncation(default), deletion, transaction (in future development)
  * For mongo:
    * truncation - default
  * For redis:
    * truncation - default

When using 'transaction' strategy you need to call Start() before the tests because it needs to know to open up a transaction.

Available options for truncation strategies:
  
  * except: ["people", "animals"] - deletes all tables except 'people' and 'animals'
  * only: ['people', 'animals'] - deletes only 'people' and 'animals' tables
