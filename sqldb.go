package clean_like_gopher

// import "database/sql"

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

// Clean with Mysql adapter
func (m *Mysql) Clean() {}

// Clean with Mysql adapter - transaction strategy
func (m *Mysql) CleanWithTransaction() {}

// Clean with Mysql adapter - truncation strategy
func (m *Mysql) CleanWithTruncation() {}

// Clean with Mysql adapter - deletion strategy
func (m *Mysql) CleanWithDeletion() {}

// For debug purposes
func (m Mysql) String() string{
  return "Mysql adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy
}
