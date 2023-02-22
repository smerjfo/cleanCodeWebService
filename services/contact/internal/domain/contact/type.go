package contact

import (
	"errors"
	"fmt"
	"regexp"
)

const pattenPhone = "[0-9]{11}"

type Contact struct {
	id          string
	fullName    string
	phoneNumber string
}

func (c *Contact) getName() string {
	return c.fullName
}

func New(id string, firstName string, middleName string, surName string, phoneNumber string) (*Contact, error) {
	match, err := regexp.Match(pattenPhone, []byte(phoneNumber))
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("phone number must consists of digits only")
	}
	return &Contact{
		id:          id,
		fullName:    fmt.Sprintf("%s %s %s", surName, firstName, middleName),
		phoneNumber: phoneNumber,
	}, nil
}
