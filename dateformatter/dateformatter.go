package dateformatter

import "time"

//GetDate returns any date matching the format format in the string data
func GetDate(data string, format string) time.Time {
	n := 0
	var parsed time.Time
	var err error
	for n < 1+len(data)-len(format) {
		parsed, err = time.Parse(format, data[n:n+len(format)])
		if err == nil {
			break
		}
		n++
	}
	return parsed
}
