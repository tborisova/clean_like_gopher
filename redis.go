package clean_like_gopher

// import "gopkg.in/redis.v2"

/*
  Redis type fields:
   name - contains the DB name
   strategy - contains the selected strategy for cleaning the DB
   options - options for additional info - [except, only]
*/
type Redis struct {
	Name    string
	Stategy string
	Options map[string][]string
}

// Clean with Redis adapter
func (m *Redis) Clean(options map[string][]string) {}
func (m *Redis) Close()                            {}

// Clean with Redis adapter - truncation strategy
func (m *Redis) CleanWithTruncation() {}

func NewRedisCleaningGopher(options map[string]string) *Redis {
	return &Redis{Name: options["dbName"]}
}

// For debug purposes
func (m Redis) String() string {
	return "Redis adapter, " + "database name: " + m.Name + ", Stategy: " + m.Stategy
}
