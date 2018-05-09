package jsondoc

import (
	"encoding/json"
	ioutil "io/ioutil"
	"github.com/spf13/cast"
	"github.com/metakeule/fmtdate"
	"time"

)

type JsonDOC struct {
	values map[string] interface{}
}

//NewJson() Return new JsonDoc pointer
func NewJson() *JsonDOC {
	return &JsonDOC{values: make(map[string]interface{})}
}

//LoadFromFile( path string) error
// Load json from file
func (j *JsonDOC)LoadFromFile( path string) error{
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &j.values)
	if err != nil {
		return err
	}
	return nil
}

//LoadFromString(str string) error
// Create a json struct using a json formated string
func (j *JsonDOC) LoadFromString(str string) error {
	err := json.Unmarshal([]byte(str), &j.values)
	if err != nil {
		return err
	}

	return nil
}

// Save(path string) error
// Save a json file based on structure
func (j *JsonDOC) Save(path string) error {
	b, err := json.MarshalIndent(&j.values, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0644)
}

//AddValue(key string, value interface{})
//Add a value to root of the json
func (j *JsonDOC) AddValue(key string, value interface{}) {
	j.values[key] = value
}

//AddJson(key string, doc *JsonDOC)
//Add a json object into root json
func (j *JsonDOC) AddJson(key string, doc *JsonDOC) {
	j.values[key] = doc.values
}

//AddJsonSlice(key string, docs []*JsonDOC)
//Add an array of json objects to the root json
func (j *JsonDOC) AddJsonSlice(key string, docs []*JsonDOC) {
	var ar []map[string]interface{}
	for _,i := range docs {
		ar = append(ar,i.values)
	}
	j.values[key] = ar
}

//GetInterface(key string) interface{}
// Get a interface using this key
func (j *JsonDOC) GetInterface(key string) interface{} {
	return j.values[key]
}

//GetInterfaceSlice(key string) []interface{}
// Get an array of interfaces
func (j *JsonDOC) GetInterfaceSlice(key string) []interface{} {
	return cast.ToSlice(j.values[key])
}

//GetInt(key string) int
// Get a int, if error or not found return 0
func (j *JsonDOC) GetInt(key string) int {
	return cast.ToInt(j.values[key])
}

//GetInt64(key string) int64
// Get a int64, if error or not found return 0
func (j *JsonDOC) GetInt64(key string) int64 {
	return cast.ToInt64(j.values[key])
}

//GetIntSlice(key string) []int
//Get an array of int
func (j *JsonDOC) GetIntSlice(key string) []int {
	return cast.ToIntSlice(j.values[key])
}

//GetFloat64(key string) float64
//Get a float64, if error or not found return 0
func (j *JsonDOC) GetFloat64(key string) float64 {
	return cast.ToFloat64(j.values[key])
}

//GetBool(key string) bool
// Get boolean value
func (j *JsonDOC) GetBool(key string) bool {
	return cast.ToBool(j.values[key])
}

//GetBoolSlice(key string) []bool
// Get an array of boolean
func (j *JsonDOC) GetBoolSlice(key string) []bool {
	return cast.ToBoolSlice(j.values[key])
}

// GetDate(format, key string)
//Get a time using a format

////Placeholders for format
//M    - month (1)
//MM   - month (01)
//MMM  - month (Jan)
//MMMM - month (January)
//D    - day (2)
//DD   - day (02)
//DDD  - day (Mon)
//DDDD - day (Monday)
//YY   - year (06)
//YYYY - year (2006)
//hh   - hours (15)
//mm   - minutes (04)
//ss   - seconds (05)
//
//AM/PM hours: 'h' followed by optional 'mm' and 'ss' followed by 'pm', e.g.
//
//hpm        - hours (03PM)
//h:mmpm     - hours:minutes (03:04PM)
//h:mm:sspm  - hours:minutes:seconds (03:04:05PM)
//
//Time zones: a time format followed by 'ZZZZ', 'ZZZ' or 'ZZ', e.g.
//
//hh:mm:ss ZZZZ (16:05:06 +0100)
//hh:mm:ss ZZZ  (16:05:06 CET)
//hh:mm:ss ZZ   (16:05:06 +01:00)
func (j *JsonDOC) GetDate(format, key string) (time.Time, error) {
	str := cast.ToString(j.values[key])
	return fmtdate.Parse(format,str)
}

//GetTime(key string) (time.Time)
// Return Time structure
func (j *JsonDOC) GetTime(key string) (time.Time) {
	return cast.ToTime(j.values[key])
}

//GetDuration(key string) time.Duration
// Returna a duration
func (j *JsonDOC) GetDuration(key string) time.Duration {
	return cast.ToDuration(j.values[key])
}

//GetString(key string) string
// Return a string
func (j *JsonDOC) GetString(key string) string {
	return cast.ToString(j.values[key])
}

//GetStringSlice(key string) []string
//Return an array of strings
func (j *JsonDOC) GetStringSlice(key string) []string {
	return cast.ToStringSlice(j.values[key])
}

//GetJson(key string) *JsonDOC
// Return object json from root json
func (j *JsonDOC) GetJson(key string) *JsonDOC {
	js := cast.ToStringMap(j.values[key])
	return &JsonDOC{values:js}
}

//GetJsonSlice(key string) []*JsonDOC
// Return an array of json objects from root json
func (j *JsonDOC) GetJsonSlice(key string) []*JsonDOC{
	var ret []*JsonDOC
	ar := cast.ToSlice(j.values[key])
	for _, i := range ar {
		v := cast.ToStringMap(i)
		ret = append(ret,&JsonDOC{values:v})
	}
	return ret
}
