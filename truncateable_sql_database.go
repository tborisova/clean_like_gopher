package clean_like_gopher

import (
	"database/sql"
	"errors"
	"fmt"
)

type TruncateableSqlDatabase interface {
	TableNames() []string
	DB() *sql.DB
}

type TruncateableSqlCleaner struct {
	adapter    string
	connection TruncateableSqlDatabase
	dbName     string
}

var UnsupportedDriver = errors.New("Unsupported database driver.")

func NewTruncateableSqlCleaningGopher(adapter string, options map[string]string) (*TruncateableSqlCleaner, error) {
	cleaner := TruncateableSqlCleaner{
		adapter: adapter,
		dbName:  options["dbName"],
	}
	var err error

	if adapter == "mysql" {
		cleaner.connection, err = NewMysqlConnection(options)
	} else if adapter == "postgres" {
		cleaner.connection, err = NewPostgresConnection(options)
	} else {
		err = UnsupportedDriver
	}

	if err != nil {
		return nil, err
	} else {
		return &cleaner, nil
	}
}

func (tsc *TruncateableSqlCleaner) Clean(options map[string][]string) {
	strategy := SelectStrategy(options)
	if strategy == "truncation" {
		err := tsc.CleanWithStatment(options, "TRUNCATE ")
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := tsc.CleanWithStatment(options, "DELETE FROM ")
		if err != nil {
			panic(err.Error())
		}
	}
}

func (tsc *TruncateableSqlCleaner) CleanWithStatment(options map[string][]string, stm string) error {
	tablesNames := tsc.connection.TableNames()

	for _, table := range tablesNames {
		if CollectionCanBeDeleted(table, options) {
			statement, err := tsc.connection.DB().Prepare(stm + table)

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

// closes the connection to the DB
func (tsc *TruncateableSqlCleaner) Close() {
	tsc.connection.DB().Close()
}

// For debug purposes
func (tsc *TruncateableSqlCleaner) String() string {
	return fmt.Sprintf("%s adapter, database name: %s", tsc.adapter, tsc.dbName)
}
