package models

import (
	"testing"
)

func TestCreateAuthorRequestValidation(t *testing.T) {
	tests := []struct {
		name        string
		payload     CreateAuthorRequest
		expectedErr error
	}{
		{
			name: "valid request",
			payload: CreateAuthorRequest{
				Name: "John Doe",
			},
			expectedErr: nil,
		},
		{
			name: "missing name",
			payload: CreateAuthorRequest{
				Name: "",
			},
			expectedErr: ErrMissingName,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.payload.Validate()
			if err != test.expectedErr {
				t.Errorf("expected error: %v, got error: %v", test.expectedErr, err)
			}
		})
	}
}

var ErrMissingName = ValidationError("missing name")

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

func (r *CreateAuthorRequest) Validate() error {
	if r.Name == "" {
		return ErrMissingName
	}
	return nil
}
