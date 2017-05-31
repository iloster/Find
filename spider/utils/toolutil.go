package utils

import (
	"time"
	"strings"
)

func ParseTime(formatted string) (time.Time, error) {
	var layouts = [...]string{
		"Mon, _2 Jan 2006 15:04:05 CCT",
		"Mon, _2 Jan 2006 15:04:05 +0000",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		"Mon, 2, Jan 2006 15:4",
		"02 Jan 2006 15:04:05 CCT",
	}
	var t time.Time
	var err error
	formatted = strings.TrimSpace(formatted)
	loc, _ := time.LoadLocation("Local")
	for _, layout := range layouts {
		t, err = time.ParseInLocation(layout, formatted,loc)
		if !t.IsZero() {
			break
		}
	}
	return t, err
}