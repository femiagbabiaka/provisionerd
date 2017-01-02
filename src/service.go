package main

import (
	"errors"
)

// Provisionerd is an exported type that
// contains the various methods our HTTP handlers
// will need to provision various pieces of infrastructure.
type Provisionerd interface {
	AddVirtualMailer(virtualMailer) (string, error)
	RemoveVirtualMailer(virtualMailer) (string, error)
	GetVirtualMailer(id int) (virtualMailer, error)
}

type provisionerd struct{}

func (provisionerd) AddVirtualMailer(vm virtualMailer) (virtualMailer, error) {
	vm, err := vm.CreateMailer()
	
	if err != nil {
		return vm, ErrCreateMailer
	}
	
	return vm, nil
}

func (provisionerd) RemoveVirtualMailer(id int) (bool, error) {
	status, err := DeleteMailer(id)
	
	if err != nil {
		return false, ErrDeleteMailer
	}
	
	return status, nil
}

// ErrCreateMailer An error representing issues with mailer creation.
var ErrCreateMailer = errors.New("Couldn't create the mailer.")
// ErrDeleteMailer An error representing issues with mailer deletion.
var ErrDeleteMailer = errors.New("Couldn't delete the mailer.")
