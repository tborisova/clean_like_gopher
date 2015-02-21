package clean_like_gopher

import (
	"fmt"
	"gopkg.in/mgo.v2"
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

var mongoStartOptions = map[string]string{"host": "localhost", "dbName": "test", "port": "27017"}

func makeDbDirty(session *mgo.Session) {
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

func CountRows(dbName, collectionName string, session *mgo.Session) int {
	collection := session.DB(dbName).C(collectionName)
	count, _ := collection.Count()

	return count
}

func TestNewCleaningGopherMongoWithIncorrectConnection(t *testing.T) {
	dbIncorrectOptions := map[string]string{"host": "localhost", "dbName": "test", "port": "27010"}
	_, err := NewCleaningGopher("mongo", dbIncorrectOptions)

	if err == nil {
		t.Errorf("Expected error since the host does not exists")
	}
}

func TestMongoCleanAll(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	makeDbDirty(session)

	m.Clean(nil)
	m.Close()

	countOfPeople := CountRows("test", "people", session)
	countOfAnimals := CountRows("test", "animals", session)

	if countOfAnimals != 0 || countOfPeople != 0 {
		t.Errorf("Expected db to be empty")
	}
}

func TestMongoCleanOnly(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	makeDbDirty(session)

	options := make(map[string][]string)
	options["only"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	countOfPeople := CountRows("test", "people", session)
	countOfAnimals := CountRows("test", "animals", session)

	if countOfPeople == 0 {
		t.Errorf("Expected people collection to not be deleted")
	}

	if countOfAnimals != 0 {
		t.Errorf("Expected animals to be deleted")
	}
}

func TestMongoCleanExcept(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	makeDbDirty(session)

	options := make(map[string][]string)
	options["except"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	countOfAnimals := CountRows("test", "animals", session)
	countOfPeople := CountRows("test", "people", session)

	if countOfAnimals == 0 {
		t.Errorf("Expected animals to not be deleted")
	}
	if countOfPeople != 0 {
		t.Errorf("Expected people to be deleted")
	}
}

func TestNewCleaningGopherMongoInvalidOptionsAreIgnored(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	makeDbDirty(session)

	options := make(map[string][]string)
	options["invalid"] = []string{"animals"}
	m.Clean(options)
	m.Close()

	countOfAnimals := CountRows("test", "animals", session)
	countOfPeople := CountRows("test", "people", session)

	if countOfAnimals != 0 {
		t.Errorf("Expected animals to be deleted")
	}
	if countOfPeople != 0 {
		t.Errorf("Expected people to be deleted")
	}
}

func TestNewCleaningGopherMongoWithMissingDbName(t *testing.T) {
	dbIncorrectOptions := map[string]string{"host": "localhost", "port": "27017"}
	_, err := NewCleaningGopher("mongo", dbIncorrectOptions)

	if err == nil {
		t.Errorf("Expected error since the db is not specified")
	}
}

func TestMongoCleanTwice(t *testing.T) {
	m, _ := NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	makeDbDirty(session)

	m.Clean(nil)

	countOfPeople := CountRows("test", "people", session)
	countOfAnimals := CountRows("test", "animals", session)

	if countOfAnimals != 0 || countOfPeople != 0 {
		t.Errorf("Expected db to be empty")
	}

	makeDbDirty(session)

	m.Clean(nil)

	countOfPeople = CountRows("test", "people", session)
	countOfAnimals = CountRows("test", "animals", session)

	if countOfAnimals != 0 || countOfPeople != 0 {
		t.Errorf("Expected db to be empty")
	}

	m.Close()
}
func TestNewCleaningGopherMongoWithIncorrectStrategy(t *testing.T) {}
