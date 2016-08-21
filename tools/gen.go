package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

func _panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func execTemplate(args interface{}) string {
	funcMap := template.FuncMap{
		"upper": func(str string) string {
			return strings.ToUpper(str[0:1]) + str[1:]
		},
	}

	t, err := template.New("gen.tmpl").Funcs(funcMap).ParseFiles("./tools/gen.tmpl")
	_panic(err)

	b := bytes.NewBuffer(nil)
	w := bufio.NewWriter(b)

	err = t.Execute(w, args)
	_panic(err)

	w.Flush()

	return string(b.Bytes())
}

func main() {
	types := []string{
		"string",
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"rune",
		"byte",
	}
	fmt.Println(execTemplate(types))
}
