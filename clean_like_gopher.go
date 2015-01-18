package clean_like_gopher

type Generic interface{
  Clean()
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
