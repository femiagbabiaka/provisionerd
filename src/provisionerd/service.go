package main

import (
	"errors"
)

// Provisionerd is an exported type that
// contains the various methods our HTTP handlers
// will need to provision various pieces of infrastructure.
type Provisionerd interface {
	AddVirtualMailer(VirtualMailer) (VirtualMailer, error)
	RemoveVirtualMailer(int) (bool, error)
}

type provisionerd struct{}

func (provisionerd) AddVirtualMailer(vm VirtualMailer) (VirtualMailer, error) {
	vm, err := vm.CreateMailer()
	
	if err != nil {
		return VirtualMailer{}, ErrCreateMailer
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
