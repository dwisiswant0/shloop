package main

import (
	"bytes"
	"text/template"
)

func parse(command string) string {
	var tpl bytes.Buffer

	funcMap := template.FuncMap{
		"i":    increment,
		"str":  rstr,
		"int":  rint,
		"uuid": uuid,
	}

	t, err := template.New("tpl").Funcs(funcMap).Parse(command)
	if err != nil {
		return ""
	}

	if err := t.Execute(&tpl, command); err != nil {
		return ""
	}

	return tpl.String()
}
