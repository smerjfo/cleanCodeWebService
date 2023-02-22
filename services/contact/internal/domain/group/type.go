package group

import (
	"errors"
	"regexp"
)

const patternName = ".{0,250}"

type Group struct {
	id   string
	name string
}

func New(id string, name string) (*Group, error) {
	match, err := regexp.Match(patternName, []byte(name))
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("maximum length of the name is 250 characters")
	}
	return &Group{id: id, name: name}, nil
}
