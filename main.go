package main

import (
	"github.com/chenyingzhou/octopus-go/claw"
)

func main() {
	c := claw.Container[1]
	values := make(map[string][]string)
	values["id"] = make([]string, 0)
	values["id"] = append(values["id"], "10001")
	sf := claw.SourceFilter{
		Type:   "AND",
		Values: values,
	}
	data := make(map[string][]map[string]string)
	c.SourceTree.Fetch(sf, &data)
}
