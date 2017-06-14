package types

import "errors"

// A Mutator is a mutator specification.
type Mutator struct {
	// Name is the unique identifier for a mutator.
	Name string `json:"name"`

	// Command is the command to be executed.
	Command string `json:"command"`

	// Timeout is the command execution timeout in seconds.
	Timeout int `json:"timeout"`

	// Env is a list of environment variables to use with command execution
	Env []string `json:"environment,omitempty"`

	// Organization specifies the organization to which the mutator belongs.
	Organization string
}

// Validate returns an error if the mutator does not pass validation tests.
func (m *Mutator) Validate() error {
	err := validateName(m.Name)
	if err != nil {
		return errors.New("mutator name " + err.Error())
	}

	if m.Command == "" {
		return errors.New("mutator command must be set")
	}

	if m.Organization == "" {
		return errors.New("mutator organization must be set")
	}

	return nil
}

// FixtureMutator returns a Mutator fixture for testing.
func FixtureMutator(name string) *Mutator {
	return &Mutator{
		Name:         name,
		Command:      "command",
		Organization: "default",
	}
}