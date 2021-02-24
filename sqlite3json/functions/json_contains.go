package functions

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pauloavelar/go-sqlite3-json/sqlite3json/utils"
)

const fnNameJsonContains = "json_contains"

func jsonContains(target, candidate interface{}) (bool, error) {
	if utils.IsNil(target) {
		return utils.IsNil(candidate), nil
	}

	if utils.IsNil(candidate) {
		return false, nil
	}

	parsedTarget, err := parseJSON(target, "target")
	if err != nil {
		return false, err
	}

	parsedCandidate, err := parseJSON(candidate, "candidate")
	if err != nil {
		return false, err
	}

	switch typedTarget := parsedTarget.(type) {
	case []interface{}:
		return utils.IsValueInArray(typedTarget, parsedCandidate), nil
	default:
		return false, newError("unsupported target type - incomplete implementation")
	}
}

func parseJSON(value interface{}, field string) (res interface{}, err error) {
	if reflect.TypeOf(value).Kind() != reflect.String {
		return false, newError("%s is not a serialized JSON: %v", field, value)
	}

	err = json.Unmarshal([]byte(value.(string)), &res)
	if err != nil {
		return false, newError("%s is not a valid JSON: %v - %s", field, value, err.Error())
	}

	return res, nil
}

func newError(message string, args ...interface{}) error {
	return fmt.Errorf(fnNameJsonContains+" - "+message, args...)
}
