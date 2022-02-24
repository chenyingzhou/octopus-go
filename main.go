package main

import (
	"github.com/chenyingzhou/octopus-go/claw"
)

func main() {
	c := claw.Container[1]
	values := make(map[string][]string)
	values["name"] = make([]string, 0)
	values["name"] = append(values["name"], "16")
	sf := claw.SourceFilter{
		Type:   "AND",
		Values: values,
	}
	data := new(map[string][]map[string]string)
	c.SourceTree.Fetch(sf, data)
}
