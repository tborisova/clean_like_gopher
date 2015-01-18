# Clean like gopher

Clean like gopher is inspired from [database_cleaner](https://github.com/DatabaseCleaner/database_cleaner). The purpose of clean
like gopher is to ensure that the DB is clean between tests.

## Supported drivers
   * [database/sql](http://golang.org/pkg/database/sql/) 
   * [redis](https://github.com/go-redis/redis)
   * [mongo](https://labix.org/mgo)

## Install:
```
  go get "github.com/tborisova/clean_like_gopher"
```
## Usage

```go
  import 'github/tborisova/clean_like_gopher'
  ...

  options := make(map[string][]string)
  options["only"] = []string{"people"}
  m := clean_like_gopher.NewCleaningGopher("test", "mongo", "truncation", options) // clean collection 'test' using mongo driver and truncation strategy
  // m.Start() - only for transaction strategy
  ...
  dirty db
  ...
  m.Clean()
```

Availabe strategies:

  * For mysql:
    * truncation, deletion, transaction
  * For mongo:
    * truncation
  * For redis:
    * truncation

When using 'transaction' strategy you need to call Start() before the tests because it needs to know to open up a transaction.

Available options for truncation strategies:
  
  * except: ["people", "animals"] - deletes all tables except 'people' and 'animals'
  * only: ['people', 'animals'] - deletes only 'people' and 'animals' tables

