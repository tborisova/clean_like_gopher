package clean_like_gopher

import (
	"fmt"
	"strings"
)

type Generic interface {
	Clean(options map[string][]string)
	Close()
}

type GopherError struct {
	Message string
}

func (e *GopherError) Error() string {
	return e.Message
}

// Creates new Cleaner based on the chosen adapter
func NewCleaningGopher(adapter string, options map[string]string) (Generic, error) {
	if adapter == "mongo" {
		cleaner, err := NewMongoCleaningGopher(options)
		return cleaner, err
	}

	if adapter == "mysql" {
		cleaner, err := NewMysqlCleaningGopher(options)
		return cleaner, err
	}

	if adapter == "redis" {
		cleaner := NewRedisCleaningGopher(options)
		return cleaner, nil
	}

	return nil, &GopherError{Message: fmt.Sprintf("Adapter %s is not supported!", adapter)}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// checks if collection can be deleted by selected options
func CollectionCanBeDeleted(name string, options map[string][]string) bool {
	if strings.Contains(name, "system") {
		return false
	}

	if len(options) == 0 {
		return true
	}

	if len(options["only"]) == 0 && len(options["except"]) != 0 {
		return !stringInSlice(name, options["except"])
	}

	if len(options["except"]) == 0 && len(options["only"]) != 0 {
		return stringInSlice(name, options["only"])
	}

	return true
}

// returns specified strategy 
func SelectStrategy(options map[string][]string) string {
	if len(options) != 0 && len(options["stategy"]) != 0 {
		return options["strategy"][0]
	} else {
		return "truncation"
	}
}
