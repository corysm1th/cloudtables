package cloudtables

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// Valid States

	SyncInProgress = "SyncInProgress"
	SyncComplete   = "SyncComplete"

	// Valid Providers

	AWS = "AWS"
)

// State is a state machine that tracks the status of
// all accounts being managed by CloudTables
type State struct {
	Accounts []Account
}

// Account holds a cloud provider account tied to the provider's API endpoint
type Account struct {
	Provider string
	Name     string
	State    string
}

// NewState returns a pointer to a State
func NewState() *State {
	s := &State{}
	return s
}

// AddAccount adds an account to the state machine
func (s *State) AddAccount(provider, name string) error {
	a := Account{State: "stale"}
	switch provider {
	case AWS:
		a.Provider = AWS
	default:
		return errors.New(fmt.Sprintf("Unknown provider: %s", provider))
	}
	if name == "" {
		return errors.New("Name must not be empty")
	}
	a.Name = name
	return nil
}

// SetState sets the state of an account.  Valid states are "in_progress", and "complete"
func (s *State) SetState(provider, name, state string) error {
	// Check for empty values
	if provider == "" {
		return errors.New("Provider required")
	}
	if name == "" {
		return errors.New("Name required")
	}
	if state == "" {
		return errors.New("State required")
	}

	// Set the state
	for _, account := range s.Accounts {
		if account.Name == name && account.Provider == provider {
			switch state {
			case SyncInProgress:
				account.State = SyncInProgress
				return nil
			case SyncComplete:
				account.State = SyncComplete
				return nil
			default:
				return errors.New(fmt.Sprintf("Unknown state: %s", state))
			}
		}
	}
	return errors.New("Account not found")
}

// GetState returns the state of all accounts
func (s *State) GetState() []Account {
	return s.Accounts
}
