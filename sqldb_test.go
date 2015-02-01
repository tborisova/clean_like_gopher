package clean_like_gopher

import "testing"

func TestMysqlNewCleaningGopherWithNonExistingDB(t *testing.T) {}

func TestMysqlNewCleaningGopherWithNonExistingAdapter(t *testing.T) {}
func TestMysqlNewCleaningGopherWithIncorrectOptions(t *testing.T)   {}
func TestMysqlNewCleaningGopherWithIncorrectStrategy(t *testing.T)  {}

func TestMysqlCleanOnly(t *testing.T)            {}
func TestMysqlCleanExcept(t *testing.T)          {}
func TestMysqlCleanAll(t *testing.T)             {}
func TestMysqlCleanWithTransaction(t *testing.T) {}
func TestMysqlCleanWithTruncation(t *testing.T)  {}
func TestMysqlCleanWithDeletion(t *testing.T)    {}
