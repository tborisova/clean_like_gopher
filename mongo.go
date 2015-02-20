package clean_like_gopher

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
  Mongo type fields:
   name - contains the DB name
   session - new connection to the db
*/
type Mongo struct {
	session *mgo.Session
	DbName  string
}

func NewMongoCleaningGopher(options map[string]string) (*Mongo, error) {
	host, ok := options["host"]
	if !ok{
		host = "@localhost"
	} else {
		host = host
	}

	port, ok := options["port"]
	if !ok{
		port = ""
	} else {
		port = ":" + port
	}
	
	dbName, ok := options["dbName"]
	if !ok{
		return nil, &GopherError{Message: "missing db name!"}
	} else {
		dbName = dbName
	}

	username, ok := options["username"]
	if !ok{
		username = ""
	} else{
		username = username
	}

	pass, ok := options["password"]
	if !ok{
		pass = ""
	} else {
		pass = ":" + pass + "@"
	}

	s, err := mgo.DialWithTimeout(username + pass + host + port, time.Duration(5)*time.Second)
	if err != nil {
		return nil, err
	} else {
		return &Mongo{DbName: dbName, session: s}, nil
	}
}

// Connect to DB
func (m *Mongo) Start() {}

// Clean with Mongo adapter
func (m *Mongo) Clean(options map[string][]string) {
	db := m.session.DB(m.DbName)
	collections, _ := db.CollectionNames()

	for _, collection_name := range collections {
		if CollectionCanBeDeleted(collection_name, options) {
			db.C(collection_name).RemoveAll(bson.M{})
		}
	}
}

// Clean with Mongo adapter - truncation strategy
func (m *Mongo) CleanWithTruncation() {}

// For debug purposes
func (m Mongo) String() string {
	return "Mongo adapter"
}

func (m Mongo) Close(){
	m.session.Close()
}

