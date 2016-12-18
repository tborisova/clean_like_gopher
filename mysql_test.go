package clean_like_gopher

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

const MyUser = "root"
const MyPw = ""
const MyDbName = "golangtest"

var mysqlStartOptions = map[string]string{"username": MyUser, "password": MyPw, "dbName": MyDbName}

func makeMysqlDbDirty(db *sql.DB) {

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

func TestMysqlCleanOnly(t *testing.T) {
	m, _ := NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MyUser, MyPw, MyDbName))
	defer db.Close()

	makeMysqlDbDirty(db)

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

func TestMysqlCleanExcept(t *testing.T) {
	m, _ := NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MyUser, MyPw, MyDbName))
	defer db.Close()

	makeMysqlDbDirty(db)

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

func TestMysqlCleanAll(t *testing.T) {
	m, _ := NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MyUser, MyPw, MyDbName))
	defer db.Close()

	makeMysqlDbDirty(db)

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

func TestMysqlNewCleaningGopherWithIncorrectDbOptions(t *testing.T) {
	StartOptions := map[string]string{"dbName": MyDbName}

	_, err := NewCleaningGopher("mysql", StartOptions)

	if err == nil {
		t.Errorf("Expected error for missing username!")
	}

	StartOptions = map[string]string{"username": "root"}

	_, err = NewCleaningGopher("mysql", StartOptions)

	if err == nil {
		t.Errorf("Expected error for missing db name!")
	}
}

func TestMysqlCleanShouldWorkTwice(t *testing.T) {
	m, _ := NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MyUser, MyPw, MyDbName))
	defer db.Close()

	makeMysqlDbDirty(db)

	m.Clean(nil)

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

	makeMysqlDbDirty(db)
	m.Clean(nil)

	row = db.QueryRow("SELECT count(id) as count FROM people")
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

	m.Close()

}

func TestMysqlNewCleaningGopherWithIncorrectStrategy(t *testing.T) {} //mysql uses truncation by default
func TestMysqlCleanWithTruncation(t *testing.T)                    {} // not sure how to test it
func TestMysqlCleanWithDeletion(t *testing.T)                      {} // not sure how to test it
