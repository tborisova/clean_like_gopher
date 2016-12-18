package clean_like_gopher

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

const PgUser = "root"
const PgPw = ""
const PgDbName = "golangtest"

var pgStartOptions = map[string]string{"host_port": "localhost:5432", "username": PgUser, "password": PgPw, "dbName": PgDbName, "connParams": "sslmode=disable"}

func makePostgresDbDirty(db *sql.DB) {

	stmtIns, err := db.Prepare("INSERT INTO users (name, age) VALUES('peter', 12)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmtIns.Exec()

	defer stmtIns.Close()

	stmtIns, err = db.Prepare("INSERT INTO people (name, age) VALUES('ivan', 12)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmtIns.Exec()

	stmtIns, err = db.Prepare("INSERT INTO animals(name) VALUES('dog')")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmtIns.Exec()
}

func TestPostgresCleanOnly(t *testing.T) {
	m, _ := NewCleaningGopher("postgres", pgStartOptions)
	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", PgUser, PgPw, PgDbName))
	defer db.Close()

	makePostgresDbDirty(db)

	options := make(map[string][]string)
	options["only"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	var count int
	row := db.QueryRow("SELECT count(*) as count FROM people")
	_ = row.Scan(&count)

	if count == 0 {
		t.Errorf("Expected people collection to not be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM users")
	_ = row.Scan(&count)

	if count == 0 {
		t.Errorf("Expected users collection to not be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM animals")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected animals to be deleted")
	}
}

func TestPostgresCleanExcept(t *testing.T) {
	m, _ := NewCleaningGopher("postgres", pgStartOptions)
	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", PgUser, PgPw, PgDbName))
	defer db.Close()

	makePostgresDbDirty(db)

	options := make(map[string][]string)
	options["except"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	var count int
	row := db.QueryRow("SELECT count(*) as count FROM people")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected people collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM users")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected users collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM animals")
	_ = row.Scan(&count)

	if count == 0 {
		t.Errorf("Expected animals to not be deleted")
	}
}

func TestPostgresCleanAll(t *testing.T) {
	m, _ := NewCleaningGopher("postgres", pgStartOptions)
	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", PgUser, PgPw, PgDbName))
	defer db.Close()

	makePostgresDbDirty(db)

	m.Clean(nil)
	m.Close()

	var count int
	row := db.QueryRow("SELECT count(*) as count FROM people")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected people collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM users")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected users collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM animals")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected animals to be deleted")
	}
}

func TestPostgresNewCleaningGopherWithIncorrectDbOptions(t *testing.T) {
	StartOptions := map[string]string{"dbName": DbName}

	_, err := NewCleaningGopher("postgres", StartOptions)

	if err == nil {
		t.Errorf("Expected error for missing username!")
	}

	StartOptions = map[string]string{"username": "root"}

	_, err = NewCleaningGopher("postgres", StartOptions)

	if err == nil {
		t.Errorf("Expected error for missing db name!")
	}
}

func TestPostgresCleanShouldWorkTwice(t *testing.T) {
	m, _ := NewCleaningGopher("postgres", pgStartOptions)
	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", PgUser, PgPw, PgDbName))
	defer db.Close()

	makePostgresDbDirty(db)

	m.Clean(nil)

	var count int
	row := db.QueryRow("SELECT count(*) as count FROM people")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected people collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM users")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected users collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM animals")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected animals to be deleted")
	}

	makePostgresDbDirty(db)
	m.Clean(nil)

	row = db.QueryRow("SELECT count(*) as count FROM people")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected people collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM users")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected users collection to be deleted")
	}

	row = db.QueryRow("SELECT count(*) as count FROM animals")
	_ = row.Scan(&count)

	if count != 0 {
		t.Errorf("Expected animals to be deleted")
	}

	m.Close()

}
