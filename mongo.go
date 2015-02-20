package clean_like_gopher

import (
	"gopkg.in/mgo.v2"
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

func NewMongoCleaningGopher(name, host, port string) (*Mongo, error) {
	s, err := mgo.Dial(host + ":" + port)
	if err != nil {
		return nil, err
	} else {
		return &Mongo{DbName: name, session: s}, nil
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
