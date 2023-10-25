package gophp

import (
	"github.com/onuragtas/gophp/serialize"
	"log"
	"testing"
)

type Location struct {
	Id   int    `json:"id" php:"id"`
	Name string `json:"name" php:"name"`
}

func TestSerialize(t *testing.T) {
	var locations []Location
	locations = append(locations, Location{Id: 1, Name: "a"})
	locations = append(locations, Location{Id: 2, Name: "b"})

	test := make(map[string]interface{})
	test["locations"] = []int{1, 2, 3, 4, 5}

	marshal, err := serialize.Marshal(test)
	log.Println(string(marshal), err)
}
