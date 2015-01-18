package clean_like_gopher

import(
    // "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

/*
  Mongo type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Mongo struct{
  Name string
  Stategy string
  Options map[string][]string
}

// Clean with Mongo adapter
func (m *Mongo) Clean() {}

// Clean with Mongo adapter - truncation strategy
func (m *Mongo) CleanWithTruncation() {}

// For debug purposes
func (m Mongo) String() string{
  return "Mongo adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy + ", options: " + m.Options["only"][0]
}
