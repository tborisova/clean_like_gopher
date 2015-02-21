package clean_like_gopher

import (
	"reflect"
	"testing"
)

func TestNewCleaningGopherMongoWithNonExistingAdapter(t *testing.T) {
	options := map[string]string{"user": "root", "dbname": "dbname", "host": "localhost", "port": "27011"}

	_, err := NewCleaningGopher("gopher", options)

	if err == nil {
		t.Error("Expected error since gopher is not adapter")
	}
}

func TestWhenAdapterIsMongoThereIsMongoInstance(t *testing.T) {
	options := map[string]string{"user": "root", "dbname": "dbname", "host": "localhost", "port": "27011"}

	gopherCleaner, _ := NewCleaningGopher("mongo", options)
	mongoInst, _ := NewMongoCleaningGopher(options)
	if reflect.TypeOf(gopherCleaner) != reflect.TypeOf(mongoInst) {
		t.Error("Expected type Mongo got %s", reflect.TypeOf(gopherCleaner))
	}
}

func TestNewCleaningGopherIsMysql(t *testing.T) {
	options := map[string]string{"user": "root", "dbname": "dbname", "host": "localhost"}

	gopherCleaner, _ := NewCleaningGopher("mysql", options)
	mysqlInst, _ := NewMysqlCleaningGopher(options)
	if reflect.TypeOf(gopherCleaner) != reflect.TypeOf(mysqlInst) {
		t.Error("Expected type Mongo got %s", reflect.TypeOf(gopherCleaner))
	}
}

func TestNewCleaningGopherIsRedis(t *testing.T) {}
