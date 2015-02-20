package clean_like_gopher

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

/*
  Mysql type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Mysql struct {
  SqlDb *sql.DB
  DbName  string
}

func NewMysqlCleaningGopher(options map[string]string) (*Mysql, error) {
  hostWithPort, ok := options["host_port"]
  if !ok {
    hostWithPort = ""
  }

  username, ok := options["username"]
  if !ok{
    return nil, &GopherError{Message: "missing username!"}
  }

  password, ok := options["password"]
  if !ok{
    password = ""
  }

  protocol, ok := options["protocol"]
  if !ok{
    protocol = ""
  }

  dbName, ok := options["dbname"]
  if !ok{
    return nil, &GopherError{"missing db name!"}
  }

  db, err := sql.Open("mysql", username + ":" + password + "@" + protocol + hostWithPort + "/" + dbName)

  if err != nil {
      return nil, err
  } else {
    return &Mysql{SqlDb: db, DbName: dbName}, nil  
  }
}

// Clean with Mysql adapter
func (m *Mysql) Clean(options map[string][]string) {
  strategy := SelectStrategy(options)
  if strategy == "truncation"{
    err := m.CleanWithStatment(options, "TRUNCATE ")
    if err != nil {
      panic(err.Error())
    }
  } else {
    err := m.CleanWithStatment(options, "DELETE FROM ")      
    if err != nil {
      panic(err.Error())
    }
  }
}

func (m *Mysql) Start()                            {}

func (m *Mysql) CleanWithStatment(options map[string][]string, stm string) error {
  tablesNames := m.TableNames()

  for _, table := range tablesNames {
    if CollectionCanBeDeleted(table, options) {
      statement, err := m.SqlDb.Prepare(stm + table)
      
      if err != nil {
        return err
      }
      defer statement.Close()

      _, err = statement.Exec() 
      if err != nil {
        return err
      }      
    }
  }

  return nil
}

// For debug purposes
func (m Mysql) String() string {
  return "Mysql adapter, " + "database name: " + m.DbName
}

func (m Mysql) TableNames() []string{
  var name string
  tablesNames := make([]string, 0)
  rows, _ := m.SqlDb.Query("show tables")

  for rows.Next() {
    _ = rows.Scan(&name)
    if(len(name) > 1){
      tablesNames = append(tablesNames, name)      
    }
  }
  return tablesNames
}

func (m Mysql) Close(){
  m.SqlDb.Close()
}
