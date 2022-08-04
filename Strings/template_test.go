package strings

import (
	"testing"
)

var template *Template

const HelloFormat = `Hello ${name}
Hello ${name}`

func init() {
	template = NewTemplate(HelloFormat)
}

func TestTemplate_Execute(t *testing.T) {
	params := map[string]interface{}{
		"name": "World",
	}
	got := template.Execute(params)
	want := "Hello World\nHello World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type Params struct {
	Name string // 注意字段需以大写开头, 否则反射获取不到值
}

func TestTemplate_ExecuteAny(t *testing.T) {
	params := Params{
		"World",
	}
	got := template.ExcuteAny(params)
	want := "Hello World\nHello World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
