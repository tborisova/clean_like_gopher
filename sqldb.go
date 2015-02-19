package clean_like_gopher

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

/*
  Mysql type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Mysql struct {
  connection *driver.Conn
  DbName  string
}

func NewMysqlCleaningGopher(name string) (*Mysql, error) {
	return &Mysql{Name: name}
  db, err := sql.Open("mysql", "root:@/gatewaty_develoment")
    if err != nil {
      return nil, err
  }

  return &Mysql{connection: db, DbName: name}
}

// Clean with Mysql adapter
func (m *Mysql) Clean(options map[string][]string) {
  db, err := sql.Open("mysql", "root:@/golangtest")
  if err != nil {
      panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
  }
  defer db.Close()

  strategy = SelectStrategy(options)
  if strategy == "truncation"{
    CleanWithTruncation(db, options)
  }else{
    CleanWithDeletion(db, options)      
  }
}
func (m *Mysql) Start()                            {}

// Clean with Mysql adapter - transaction strategy
func (m *Mysql) CleanWithTransaction() {}

// Clean with Mysql adapter - truncation strategy
func (m *Mysql) CleanWithTruncation() {}

// Clean with Mysql adapter - deletion strategy
func (m *Mysql) CleanWithDeletion() {}

// For debug purposes
func (m Mysql) String() string {
	return "Mysql adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy
}
