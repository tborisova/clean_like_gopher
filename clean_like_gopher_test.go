package clean_like_gopher

import (
	"reflect"
	"testing"
)

func TestNewCleaningGopherMongoWithNonExistingAdapter(t *testing.T) {
	_, err := NewCleaningGopher("gopher", "dbname", "localhost", "27011")

	if err == nil {
		t.Error("Expected error since gopher is not adapter")
	}
}

func TestWhenAdapterIsMongoThereIsMongoInstance(t *testing.T) {
	gopher_cleaner, _ := NewCleaningGopher("mongo", "dbname", "localhost", "27011")
	mongo_inst, _ := NewMongoCleaningGopher("dbname", "localhost", "27011")
	if reflect.TypeOf(gopher_cleaner) != reflect.TypeOf(mongo_inst) {
		t.Error("Expected type Mongo got %s", reflect.TypeOf(gopher_cleaner))
	}
}

func TestNewCleaningGopherIsMysql(t *testing.T) {}

func TestNewCleaningGopherIsRedis(t *testing.T) {}
