package clean_like_gopher

type Generic interface{
  Clean()
}

type Mongo struct{
  Name string //db name
  Stategy string //truncation, transaction
  Options map[string][]string //only, except
}

type Mysql struct{
  Name string //db name
  Stategy string //truncation, transaction
  Options map[string][]string //only, except
}

type Redis struct{
  Name string //db name
  Stategy string //truncation, transaction
  Options map[string][]string //only, except
}

func NewCleaningGopher(name, adapter, st string, options map[string][]string) Generic{
  if(adapter == "mongo"){
    return &Mongo{Name: name, Stategy: st, Options: options}
  }else if(adapter == "mysql") {
    return &Mysql{Name: name, Stategy: st, Options: options}
  }else{
    return &Redis{Name: name, Stategy: st, Options: options}
  }
}

func (m *Mongo) Clean() {}

func (m *Mysql) Clean() {}

func (m *Redis) Clean() {}

func (m *Mongo) CleanWithTruncation() {}

func (m *Redis) CleanWithTruncation() {}

func (m *Mysql) CleanWithTransaction() {}

func (m *Mysql) CleanWithTruncation() {}

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