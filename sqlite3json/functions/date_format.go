package functions

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pauloavelar/go-sqlite3-json/sqlite3json/utils"
)

const fnNameDateFormat = "date_format"

func dateFormat(date, format interface{}) (sql.NullString, error) {
	if utils.IsNil(date) {
		return sql.NullString{}, nil
	}

	formatString, ok := format.(string)
	if !ok {
		return sql.NullString{}, newError(fnNameDateFormat, "format is not a valid string")
	}

	sanitizedFormat := formatString
	for specifier, replacement := range dateSpecifiers {
		sanitizedFormat = strings.ReplaceAll(sanitizedFormat, specifier, replacement)
	}


	// TODO finish date_format

	return sql.NullString{Valid: true, String: "TODO"}, nil
}

func newError(fn, message string, args ...interface{}) error {
	return fmt.Errorf(fn+" - "+message, args...)
}

// reference: Mon Jan 2 15:04:05 -0700 MST 2006

var dateSpecifiers = map[string]string{
	// weekday
	"%a": "Mon",
	"%W": "Monday",

	// year
	"%Y": "2006",
	"%y": "06",

	// month
	"%b": "Jan",
	"%c": "1",
	"%M": "January",
	"%m": "01",

	// day
	"%D": "2nd",
	"%d": "02",
	"%e": "2",

	// second fractions
	"%f": "000000",

	// hour
	"%H": "15",
	"%h": "03",
	"%I": "03",

	// minute
	"%i": "04",

	// second
	"%S": "05",
	"%s": "05",

	// misc
	"%p": "PM",
	"%r": "03:04:05PM",
	"%T": "15:04:05",
	"%%": "%",

	// unsupported
	"%j": "UNSUPPORTED", // day of year
	"%k": "UNSUPPORTED", // hour (0..23)
	"%l": "UNSUPPORTED", // hour (1..12)
	"%U": "UNSUPPORTED", // Week (00.53) - WEEK() mode 0
	"%u": "UNSUPPORTED", // Week (00.53) - WEEK() mode 1
	"%V": "UNSUPPORTED", // Week (01.53) - WEEK() mode 2
	"%v": "UNSUPPORTED", // Week (01.53) - WEEK() mode 3
	"%w": "UNSUPPORTED", // Day of the week
	"%X": "UNSUPPORTED", // Year for the week (Sunday is the first day)
	"%x": "UNSUPPORTED", // Year for the week (Monday is the first day)
}
