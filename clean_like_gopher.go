package clean_like_gopher

import(
    // "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
    // "database/sql"
    // "gopkg.in/redis.v2"
)

type Generic interface{
  Clean()
}

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

/*
  Mysql type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Mysql struct{
  Name string
  Stategy string
  Options map[string][]string
}

/*
  Redis type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Redis struct{
  Name string
  Stategy string
  Options map[string][]string
}

// Creates new Cleaner based on the chosen adapter
func NewCleaningGopher(name, adapter, st string, options map[string][]string) Generic{
  if(adapter == "mongo"){
    return &Mongo{Name: name, Stategy: st, Options: options}
  }else if(adapter == "mysql") {
    return &Mysql{Name: name, Stategy: st, Options: options}
  }else{
    return &Redis{Name: name, Stategy: st, Options: options}
  }
}

// Clean with Mongo adapter
func (m *Mongo) Clean() {}

// Clean with Mysql adapter
func (m *Mysql) Clean() {}

// Clean with Redis adapter
func (m *Redis) Clean() {}

// Clean with Mongo adapter - truncation strategy
func (m *Mongo) CleanWithTruncation() {}

// Clean with Redis adapter - truncation strategy
func (m *Redis) CleanWithTruncation() {}

// Clean with Mysql adapter - transaction strategy
func (m *Mysql) CleanWithTransaction() {}

// Clean with Mysql adapter - truncation strategy
func (m *Mysql) CleanWithTruncation() {}

// Clean with Mysql adapter - deletion strategy
func (m *Mysql) CleanWithDeletion() {}

func (m Mongo) String() string{
  return "Mongo adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy + ", options: " + m.Options["only"][0]
}

func (m Mysql) String() string{
  return "Mysql adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy
}

func (m Redis) String() string{
  return "Redis adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy
}