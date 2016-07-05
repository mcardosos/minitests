package booleangroup_test

import (
	. "Tests/body-boolean"
	"strings"
	"testing"
)

var boolClient = getBooleanClient()

func getBooleanClient() BoolGroupClient {
	c := NewBoolGroupClient()
	if strings.HasPrefix(c.BaseURI, "https") {
		c.BaseURI = "http://localhost:3000"
	} else {
		c.BaseURI += ":3000"
	}
	return c
}

func TestGetTrue(t *testing.T) {
	res, err := boolClient.GetTrue()
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if *res.Value != true {
		t.Errorf("Got %v, expecting %v\n", *res.Value, true)
	}
}

func TestPutTrue(t *testing.T) {
	_, err := boolClient.PutTrue(true)
	if err != nil {
		t.Errorf("Got %v, expecting %v\n", err, nil)
	}
}

func TestGetFalse(t *testing.T) {
	res, err := boolClient.GetFalse()
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if *res.Value != false {
		t.Errorf("Got %v, expecting %v\n", *res.Value, false)
	}
}

func TestPutFalse(t *testing.T) {
	_, err := boolClient.PutFalse(false)
	if err != nil {
		t.Errorf("Got %v, expecting %v\n", err, nil)
	}
}

func TestGetInvalidBool(t *testing.T) {
	_, err := boolClient.GetInvalid()
	if err == nil {
		t.Errorf("Got %v, expecting %v\n", err, "non nil error")
	}
}

func TestGetNullBool(t *testing.T) {
	res, err := boolClient.GetNull()
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if res.Value != nil {
		t.Errorf("Got %v, expecting %v\n", *res.Value, nil)
	}
}
