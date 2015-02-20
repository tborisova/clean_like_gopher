package clean_like_gopher

import (
	"reflect"
	"testing"
)

type Person struct {
  Name  string
  Phone string
}

type Animal struct {
  Kind      string
  Character string
}

func TestNewCleaningGopherMongoWithNonExistingAdapter(t *testing.T) {
  options := map[string]string{"user":"root", "dbname": "dbname", "host":"localhost", "port":"27011"}

	_, err := NewCleaningGopher("gopher", options)

	if err == nil {
		t.Error("Expected error since gopher is not adapter")
	}
}

func TestWhenAdapterIsMongoThereIsMongoInstance(t *testing.T) {
  options := map[string]string{"user":"root", "dbname": "dbname", "host":"localhost", "port":"27011"}

	gopher_cleaner, _ := NewCleaningGopher("mongo", options)
	mongo_inst, _ := NewMongoCleaningGopher(options)
	if reflect.TypeOf(gopher_cleaner) != reflect.TypeOf(mongo_inst) {
		t.Error("Expected type Mongo got %s", reflect.TypeOf(gopher_cleaner))
	}
}

func TestNewCleaningGopherIsMysql(t *testing.T) {}

func TestNewCleaningGopherIsRedis(t *testing.T) {}
