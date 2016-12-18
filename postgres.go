package clean_like_gopher

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

// creates new cleaner for mysql driver
func NewPostgresConnection(options map[string]string) (*Postgres, error) {
	hostWithPort, ok := options["host_port"]
	if !ok {
		hostWithPort = "localhost:5432"
	}

	username, ok := options["username"]
	if !ok {
		return nil, &GopherError{Message: "missing username!"}
	}

	password, ok := options["password"]
	if !ok {
		password = ""
	}

	connParams, ok := options["connParams"]
	if !ok {
		connParams = ""
	}

	dbName, ok := options["dbName"]
	if !ok {
		return nil, &GopherError{"missing db name!"}
	}

	connectionUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", username, password, hostWithPort, dbName, connParams)
	conn, err := sql.Open("postgres", connectionUrl)

	if err != nil {
		return nil, err
	} else {
		return &Postgres{db: conn}, nil
	}
}

// returns all table names
func (p *Postgres) TableNames() []string {
	var name string
	tablesNames := make([]string, 0)
	rows, _ := p.db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'")

	for rows.Next() {
		_ = rows.Scan(&name)
		if len(name) > 1 {
			tablesNames = append(tablesNames, name)
		}
	}
	return tablesNames
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}
