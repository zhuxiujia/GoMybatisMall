package utils

import "testing"

type TestObj struct {
	Id   string
	Name string
}

func TestGetArrayIds(t *testing.T) {
	var array = []TestObj{{Id: "1", Name: "s"}, {Id: "1", Name: "s"}}
	var ids = GetIds(array, "Id")
	for _, v := range ids {
		println(v)
	}
}
