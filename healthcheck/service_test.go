package healthcheck

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockService is a mock implementation of the Service interface for testing purposes.
type MockService struct {
}

func (ms *MockService) HealthcheckService() (Healthcheck, error) {
	// For testing, we create a Healthcheck with static values.
	check := Healthcheck{
		ServiceName: "Mock Service",
		Status:      "OK",
		Description: "Mock service is healthy",
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return check, nil
}

func TestHealthcheckService(t *testing.T) {
	// Create an instance of the actual service
	s := NewService()

	// Call the HealthcheckService method
	result, err := s.HealthcheckService()

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the returned Healthcheck has the expected ServiceName
	assert.Equal(t, "Golang Fundhub Services v3", result.ServiceName)

	// Add more assertions based on the expected values for your specific case

	// In addition, let's test the MockService
	mockService := &MockService{}
	mockResult, mockErr := mockService.HealthcheckService()

	// Assert that there is no error for the MockService
	assert.NoError(t, mockErr)

	// Assert that the returned Healthcheck from MockService has the expected ServiceName
	assert.Equal(t, "Mock Service", mockResult.ServiceName)

	// Add more assertions for the MockService based on the expected values
}
