package main

import (
	"errors"
	"go-mbv-go/pkg/api"
	"os"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	testCases := []struct {
		name             string
		environment      string
		expectedError    error
		dependencies api.API
	}{
		{
			name: "should return an error because no environment",
			environment: "",
			expectedError: errors.New("environment was not specified"),
			dependencies: api.API{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			 os.Setenv("ENVIRONMENT", test.environment)
			_, err := setupRouter(test.dependencies)

			if err.Error() != test.expectedError.Error() {
				t.Errorf("expected: %v, got: %v", err, test.expectedError)
			}
		})
	}
}

func TestSetupDatabase(t *testing.T) {
	testCases := []struct {
		name             string
		connectionString      string
		expectedError    error
	}{
		{
			name: "should return an error because no connection string",
			connectionString: "",
			expectedError: errors.New("connectionString was not specified"),
		},
		{
			name: "should not return an error",
			connectionString: "some-connection-string",
			expectedError: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv("DATABASE_URL", test.connectionString)
			_, err := setupDatabase()

			if err != test.expectedError {
				t.Errorf("expected: %v, got: %v", err, test.expectedError)
			}
		})
	}
}
