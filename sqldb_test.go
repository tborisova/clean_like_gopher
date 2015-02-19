package clean_like_gopher

import "testing"

func TestMysqlNewCleaningGopherWithIncorrectConnection(t *testing.T) {
  _, err := NewCleaningGopher("mysql", "foo", "localhost", "27010")

  if err == nil {
    t.Errorf("Expected error since the host does not exists")
  }
}

func TestMysqlNewCleaningGopherWithNonExistingAdapter(t *testing.T) {}
func TestMysqlNewCleaningGopherWithIncorrectOptions(t *testing.T)   {}
func TestMysqlNewCleaningGopherWithIncorrectStrategy(t *testing.T)  {}

func TestMysqlCleanOnly(t *testing.T)            {}
func TestMysqlCleanExcept(t *testing.T)          {}
func TestMysqlCleanAll(t *testing.T)             {}
func TestMysqlCleanWithTransaction(t *testing.T) {}
func TestMysqlCleanWithTruncation(t *testing.T)  {}
func TestMysqlCleanWithDeletion(t *testing.T)    {}
