# jsondoc
A structure in golang to manipulate json files

## Installation

As a library

```shell
go get github.com/rgobbo/jsondoc
```

## Usage

```go
package main

import (
	"github.com/rgobbo/jsondoc"
	"log"

)

func main() {
	j := jsondoc.NewJson()
	str := `{"name": "test", "age": 30}`
	err := j.LoadFromString(str)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("name=" ,j.GetString("name"))
	log.Println("age=", j.GetInt("age"))

	j2 := jsondoc.NewJson()
	j2.LoadFromFile("sample/test.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("name=" ,j2.GetString("name"))
	log.Println("count=", j2.GetInt("count"))
	log.Println("count=", j2.GetTime("date"))

}
```

## Documentation

### NewJson()
Return new JsonDoc pointer

### JsonDOC Methods

### LoadFromFile( path string) error
 Load json from file


### LoadFromString(str string) error
 Create a json struct using a json formated string


### Save(path string) error
 Save a json file based on structure


### AddValue(key string, value interface{})
 Add a value to root of the json


### AddJson(key string, doc *JsonDOC)
 Add a json object into root json


### AddJsonSlice(key string, docs []*JsonDOC)
 Add an array of json objects to the root json


### GetInterface(key string) interface{}
 Get a interface using this key


### GetInterfaceSlice(key string) []interface{}
 Get an array of interfaces


### GetInt(key string) int
 Get a int, if error or not found return 0


### GetInt64(key string) int64
 Get a int64, if error or not found return 0


### GetIntSlice(key string) []int
Get an array of int


### GetFloat64(key string) float64
 Get a float64, if error or not found return 0


### GetBool(key string) bool
 Get boolean value


### GetBoolSlice(key string) []bool
 Get an array of boolean


### GetDate(format, key string)
 Get a time using a format
 ( Dependency : "github.com/metakeule/fmtdate")
 Placeholders for format
 M    - month (1)
MM   - month (01)
MMM  - month (Jan)
MMMM - month (January)
D    - day (2)
DD   - day (02)
DDD  - day (Mon)
DDDD - day (Monday)
YY   - year (06)
YYYY - year (2006)
hh   - hours (15)
mm   - minutes (04)
ss   - seconds (05)

AM/PM hours: 'h' followed by optional 'mm' and 'ss' followed by 'pm', e.g.

hpm        - hours (03PM)
h:mmpm     - hours:minutes (03:04PM)
h:mm:sspm  - hours:minutes:seconds (03:04:05PM)

Time zones: a time format followed by 'ZZZZ', 'ZZZ' or 'ZZ', e.g.

hh:mm:ss ZZZZ (16:05:06 +0100)
hh:mm:ss ZZZ  (16:05:06 CET)
hh:mm:ss ZZ   (16:05:06 +01:00)


### GetTime(key string) (time.Time)
 Return Time structure


### GetDuration(key string) time.Duration
 Returna a duration


### GetString(key string) string
 Return a string


### GetStringSlice(key string) []string
 Return an array of strings


### GetJson(key string) *JsonDOC
 Return object json from root json


### GetJsonSlice(key string) []*JsonDOC
 Return an array of json objects from root json
