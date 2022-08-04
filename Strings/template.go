package strings

import (
	"fmt"
	"reflect"
	"strings"
)

type Template struct {
	format string
}

func NewTemplate(format string) *Template {
	return &Template{format}
}

func (t *Template) Execute(params map[string]interface{}) string {
	ret := t.format
	for key, val := range params {
		ret = strings.Replace(ret, "${"+key+"}", fmt.Sprintf("%s", val), -1)
	}
	return ret
}

func (t *Template) ExcuteAny(data any) string {
	params := ToMap(data)
	return t.Execute(params)
}

func ToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}
