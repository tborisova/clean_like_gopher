package clean_like_gopher

import(
	"fmt"
	"strings"
)

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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

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

func SelectStrategy(options map[string][]string){
	if len(options) != 0 && len(options["stategy"]) != 0 {
		return options["strategy"]
	} else {
		return "truncation"
	}
}