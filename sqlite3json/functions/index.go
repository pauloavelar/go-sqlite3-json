package functions

var (
	Available = map[string]interface{}{
		fnNameConcat:       concat,
		fnNameJsonContains: jsonContains,
	}
)
