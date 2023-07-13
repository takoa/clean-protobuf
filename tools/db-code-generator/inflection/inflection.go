package inflection

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

var defaultClient = pluralize.NewClient()

func ToSingularSnake(str string) string {
	words := strings.Split(strcase.ToSnake(str), "_")
	words[len(words)-1] = defaultClient.Singular(words[len(words)-1])

	return strings.Join(words, "_")
}

func ToPluralSnake(str string) string {
	words := strings.Split(strcase.ToSnake(str), "_")
	words[len(words)-1] = defaultClient.Plural(words[len(words)-1])

	return strings.Join(words, "_")
}
