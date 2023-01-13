package test_cmp

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

type Person struct {
	Name      string
	Age       int
	DateAdded time.Time
}

func CreatePerson(name string, age int) Person {
	return Person{
		Name:      name,
		Age:       age,
		DateAdded: time.Now(),
	}
}

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
