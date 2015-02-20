package clean_like_gopher

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"testing"
)

var	session, _ = mgo.Dial("localhost:27017")

func makeDbDirty() {
	c := session.DB("test").C("people")
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})

	c = session.DB("test").C("animals")
	err = c.Insert(&Animal{"gopher", "cute"},
		&Animal{"snake", "not cute"})
	if err != nil {
		fmt.Println(err)
	}
}

func TestNewCleaningGopherMongoWithIncorrectConnection(t *testing.T) {
	_, err := NewCleaningGopher("mongo", "foo", "localhost", "27010")

	if err == nil {
		t.Errorf("Expected error since the host does not exists")
	}
}

func TestMongoCleanAll(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", "test", "localhost", "27017")
	m.Start()

	makeDbDirty()

	m.Clean(nil)
	m.Close()

	people := session.DB("test").C("people")
	count, _ := people.Count()

	if count != 0 {
		t.Errorf("Expected db to be empty")
	}
}

func TestMongoCleanOnly(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", "test", "localhost", "27017")
	m.Start()

	makeDbDirty()

	options := make(map[string][]string)
	options["only"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	people := session.DB("test").C("people")
	count_of_peoples, _ := people.Count()
	animals := session.DB("test").C("animals")
	count_of_animals, _ := animals.Count()

	if count_of_peoples == 0 {
		t.Errorf("Expected people collection to not be deleted")
	}

	if count_of_animals != 0 {
		t.Errorf("Expected animals to be deleted")
	}
}

func TestMongoCleanExcept(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", "test", "localhost", "27017")
	m.Start()
	makeDbDirty()
	options := make(map[string][]string)
	options["except"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	people := session.DB("test").C("people")
	count_of_people, _ := people.Count()
	animals := session.DB("test").C("animals")
	count_of_animals, _ := animals.Count()

	if count_of_animals == 0 {
		t.Errorf("Expected animals to not be deleted")
	}
	if count_of_people != 0 {
		t.Errorf("Expected people to be deleted")
	}
}

func TestNewCleaningGopherMongoInvalidOptionsAreIgnored(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", "test", "localhost", "27017")
	m.Start()
	makeDbDirty()
	options := make(map[string][]string)
	options["invalid"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	people := session.DB("test").C("people")
	count_of_people, _ := people.Count()
	animals := session.DB("test").C("animals")
	count_of_animals, _ := animals.Count()

	if count_of_animals != 0 {
		t.Errorf("Expected animals to be deleted")
	}
	if count_of_people != 0 {
		t.Errorf("Expected people to be deleted")
	}
}

func TestNewCleaningGopherMongoWithIncorrectStrategy(t *testing.T) {}
