package clean_like_gopher

import "testing"

func TestNewCleaningGopherWithNonExistingDB(t *testing T){}
func TestNewCleaningGopherWithNonExistingAdapter(t *testing T){}
func TestNewCleaningGopherWithIncorrectOptions(t *testing T){}
func TestNewCleaningGopherWithIncorrectStrategy(t *testing T){}

func TestMongoCleanAll(t *testing T){}
func TestMongoCleanOnly(t *testing T){}
func TestMongoCleanExcept(t *testing T){}

func TestRedisCleanAll(t *testing T){}
func TestRedisCleanExcept(t *testing T){}
func TestRedisCleanOnly(t *testing T){}

func TestMysqlCleanOnly(t *testing T){}
func TestMysqlCleanExcept(t *testing T){}
func TestMysqlCleanAll(t *testing T){}
func TestMysqlCleanWithTransaction(t *testing T){}
func TestMysqlCleanWithTruncation(t *testing T){}
func TestMysqlCleanWithDeletion(t *testing T){}

