package clean_like_gopher

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

// creates new cleaner for mysql driver
func NewMysqlConnection(options map[string]string) (*Mysql, error) {
	hostWithPort, ok := options["host_port"]
	if !ok {
		hostWithPort = ""
	}

	username, ok := options["username"]
	if !ok {
		return nil, &GopherError{Message: "missing username!"}
	}

	password, ok := options["password"]
	if !ok {
		password = ""
	}

	protocol, ok := options["protocol"]
	if !ok {
		protocol = ""
	}

	dbName, ok := options["dbName"]
	if !ok {
		return nil, &GopherError{"missing db name!"}
	}

	conn, err := sql.Open("mysql", username+":"+password+"@"+protocol+hostWithPort+"/"+dbName)

	if err != nil {
		return nil, err
	} else {
		return &Mysql{db: conn}, nil
	}
}

// returns all table names
func (m *Mysql) TableNames() []string {
	var name string
	tablesNames := make([]string, 0)
	rows, _ := m.db.Query("show tables")

	for rows.Next() {
		_ = rows.Scan(&name)
		if len(name) > 1 {
			tablesNames = append(tablesNames, name)
		}
	}
	return tablesNames
}

func (m *Mysql) DB() *sql.DB {
	return m.db
}
