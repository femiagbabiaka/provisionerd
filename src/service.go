package main

import (
	"errors"
)

type Provisionerd interface {
	AddVirtualMailer(mailerInfo) (string, error)
	RemoveVirtualMailer(mailerInfo) (string, error)
}

type provisionerd struct{}

func (provisionerd) AddVirtualMailer(m mailerInfo) (string, error) {
}

func (provisionerd) RemoveVirtualMailer(m mailerInfo) (string, error) {
}

var ErrEmpty = errors.New("empty string")
