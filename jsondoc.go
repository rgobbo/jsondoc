package jsondoc

import (
	"encoding/json"
	ioutil "io/ioutil"
	"github.com/spf13/cast"
)

type JsonDOC struct {
	values map[string] interface{}
}

func NewJson() *JsonDOC {
	return &JsonDOC{values: make(map[string]interface{})}
}

func (j *JsonDOC)LoadFromFile( path string) error{
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, j)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonDOC) LoadFromString(str string) error {
	err := json.Unmarshal([]byte(str), j)
	if err != nil {
		return err
	}

	return nil
}

func (j *JsonDOC) Save(path string) error {
	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0644)
}

func (j *JsonDOC) AddValue(key string, value interface{}) {
	j.values[key] = value
}

func (j *JsonDOC) AddJson(key string, doc *JsonDOC) {
	j.values[key] = doc.values
}

func (j *JsonDOC) AddJsonSlice(key string, docs []*JsonDOC) {
	var ar []map[string]interface{}
	for _,i := range docs {
		ar = append(ar,i.values)
	}
	j.values[key] = ar
}

func (j *JsonDOC) GetInterface(key string) interface{} {
	return j.values[key]
}

func (j *JsonDOC) GetString(key string) string {
	return cast.ToString(j.values[key])
}

func (j *JsonDOC) GetInt(key string) int {
	return cast.ToInt(j.values[key])
}

func (j *JsonDOC) GetStringSlice(key string) []string {
	return cast.ToStringSlice(j.values[key])
}

func (j *JsonDOC) GetJson(key string) *JsonDOC {
	js := cast.ToStringMap(j.values[key])
	return &JsonDOC{values:js}
}

func (j *JsonDOC) GetJsonSlice(key string) []*JsonDOC{
	var ret []*JsonDOC
	ar := cast.ToSlice(j.values[key])
	for _, i := range ar {
		v := cast.ToStringMap(i)
		ret = append(ret,&JsonDOC{values:v})
	}
	return ret
}
