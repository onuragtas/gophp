package gophp

import (
	"github.com/onuragtas/gophp/serialize"
	"log"
	"testing"
	"time"
)

type Location struct {
	Id   int       `json:"id" php:"id"`
	Name string    `json:"name" php:"name"`
	T    time.Time `json:"t" php:"t"`
}

func TestSerialize(t *testing.T) {
	var locations []Location
	locations = append(locations, Location{Id: 1, Name: "a", T: time.Now()})
	locations = append(locations, Location{Id: 2, Name: "b", T: time.Now()})

	test := make(map[string]interface{})
	test["locations"] = []int{1, 2, 3, 4, 5}

	marshal, err := serialize.Marshal(locations)
	log.Println(string(marshal), err)
}
