package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiResponse(t *testing.T) {
	message := "Test Message"
	code := 200
	status := "success"
	data := map[string]interface{}{"key": "value"}

	response := ApiResponse(message, code, status, data)

	assert.Equal(t, message, response.Meta.Message)
	assert.Equal(t, code, response.Meta.Code)
	assert.Equal(t, status, response.Meta.Status)
	assert.Equal(t, data, response.Data)
}
