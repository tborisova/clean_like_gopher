package clean_like_gopher

import "fmt"

type Generic interface {
	Clean(options map[string][]string)
	Start()
}

type GopherError struct {
	Message string
}

func (e *GopherError) Error() string {
	return e.Message
}

// Creates new Cleaner based on the chosen adapter
func NewCleaningGopher(adapter, name, host, port string) (Generic, error) {
	if adapter == "mongo" {
		ad, err := NewMongoCleaningGopher(name, host, port)
		return ad, err
	}

	if adapter == "mysql" {
		ad := NewMysqlCleaningGopher(name)
		return ad, nil
	}

	if adapter == "redis" {
		ad := NewRedisCleaningGopher(name)
		return ad, nil
	}

	return nil, &GopherError{Message: fmt.Sprintf("Adapter %s is not supported!", adapter)}
}
