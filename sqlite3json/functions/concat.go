package functions

import (
	"fmt"
	"strings"

	"github.com/pauloavelar/go-sqlite3-json/sqlite3json/utils"
)

const fnNameConcat = "concat"

func concat(parts ...interface{}) (string, error) {
	strValues := make([]string, len(parts))

	for i := range parts {
		if utils.IsNil(parts[i]) {
			return "", fmt.Errorf(fnNameConcat+" - null param at index %d", i)
		}

		str, ok := parts[i].(string)
		if !ok {
			return "", fmt.Errorf(fnNameConcat+" - invalid param type at index %d", i)
		}

		strValues[i] = str
	}

	return strings.Join(strValues, ""), nil
}
