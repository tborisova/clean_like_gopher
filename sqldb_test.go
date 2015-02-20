package clean_like_gopher

import (
    "database/sql"
    "testing"
    _ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:@/golangtest")

func makeMysqlDbDirty() {

  stmtIns, err := db.Prepare("INSERT INTO users (name, age) VALUES(?, ?)")
  if err != nil {
      panic(err.Error())
  }
  _, err = stmtIns.Exec("peter", 12)

  defer stmtIns.Close()
  
  stmtIns, err = db.Prepare("INSERT INTO people (name, age) VALUES(?, ?)")
  if err != nil {
      panic(err.Error())
  }
  _, err = stmtIns.Exec("ivan", 12)
  
  stmtIns, err = db.Prepare("INSERT INTO animals(name) VALUES(?)")
  if err != nil {
      panic(err.Error())
  }
  _, err = stmtIns.Exec("dog")
}

func TestMysqlCleanOnly(t *testing.T)            {
  m, _ := NewCleaningGopher("mysql", "golangtest", "root", "10000")
  m.Start()

  makeMysqlDbDirty()

  options := make(map[string][]string)
  options["only"] = []string{"animals"}
  m.Clean(options)
  m.Close()

  var count int
  row := db.QueryRow("SELECT count(id) as count FROM people")
  _ = row.Scan(&count)

  if count == 0 {
    t.Errorf("Expected people collection to not be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM users")
  _ = row.Scan(&count)
  
  if count == 0 {
    t.Errorf("Expected users collection to not be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM animals")
  _ = row.Scan(&count)

  if count != 0 {
    t.Errorf("Expected animals to be deleted")
  }
}
func TestMysqlCleanExcept(t *testing.T)          {
  m, _ := NewCleaningGopher("mysql", "golangtest", "root", "10000")
  m.Start()

  makeMysqlDbDirty()

  options := make(map[string][]string)
  options["except"] = []string{"animals"}
  m.Clean(options)
  m.Close()

  var count int
  row := db.QueryRow("SELECT count(id) as count FROM people")
  _ = row.Scan(&count)

  if count != 0 {
    t.Errorf("Expected people collection to be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM users")
  _ = row.Scan(&count)
  
  if count != 0 {
    t.Errorf("Expected users collection to be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM animals")
  _ = row.Scan(&count)

  if count == 0 {
    t.Errorf("Expected animals to not be deleted")
  }
}
func TestMysqlCleanAll(t *testing.T)             {
  m, _ := NewCleaningGopher("mysql", "golangtest", "root", "10000")
  m.Start()

  makeMysqlDbDirty()

  m.Clean(nil)
  m.Close()

  var count int
  row := db.QueryRow("SELECT count(id) as count FROM people")
  _ = row.Scan(&count)

  if count != 0 {
    t.Errorf("Expected people collection to be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM users")
  _ = row.Scan(&count)
  
  if count != 0 {
    t.Errorf("Expected users collection to be deleted")
  }

  row = db.QueryRow("SELECT count(id) as count FROM animals")
  _ = row.Scan(&count)

  if count != 0 {
    t.Errorf("Expected animals to be deleted")
  }
}

func TestMysqlNewCleaningGopherWithIncorrectConnection(t *testing.T) {}
func TestMysqlNewCleaningGopherWithNonExistingAdapter(t *testing.T) {}
func TestMysqlNewCleaningGopherWithIncorrectOptions(t *testing.T)   {}
func TestMysqlNewCleaningGopherWithIncorrectStrategy(t *testing.T)  {}
func TestMysqlCleanWithTransaction(t *testing.T) {}
func TestMysqlCleanWithTruncation(t *testing.T)  {}
func TestMysqlCleanWithDeletion(t *testing.T)    {}
