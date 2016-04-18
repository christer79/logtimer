package dateformatter

import (
	"testing"
	"time"
)

type testdata struct {
	data   string
	format string
	expect time.Time
}

var testDataFormat = []testdata{

	{
		"2016-Apr-18 19:04:04 +0000",
		"2006-Jan-02 15:04:05 +0000",
		time.Date(2016, time.April, 18, 19, 04, 04, 00, time.UTC),
	},
	{
		"18/Apr/2016:19:04:04",
		"02/Jan/2006:15:04:05",
		time.Date(2016, time.April, 18, 19, 04, 04, 00, time.UTC),
	},
	{
		"10.67.6.50 - - [18/Apr/2016:19:04:04 +0000] \"GET /nexus/content/groups/public/com/jeppesen/jcms/sys-layout/maven-metadata.xml HTTP/1.1\" 200 1481 20",
		"02/Jan/2006:15:04:05",
		time.Date(2016, time.April, 18, 19, 04, 04, 00, time.UTC),
	},
}

func TestParser(t *testing.T) {
	for _, testdata := range testDataFormat {
		result := GetDate(testdata.data, testdata.format)
		if result != testdata.expect {
			t.Fatalf("Expected \"%v\" to result in: \n%v, got \n%v", testdata.data, testdata.expect, result)
		}

	}
}
