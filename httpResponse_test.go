package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetHTTPInternalServerError tests HTTP response.
func TestGetHTTPResponse(t *testing.T) {
	actual := HTTPResponse{200, "Success", nil}
	expected := GetHTTPResponse(200, "Success", nil)
	assert.Equal(t, expected, actual)

	actual = HTTPResponse{404, "Not found", 15}
	expected = GetHTTPResponse(404, "Not found", 15)
	assert.Equal(t, expected, actual)
}

// TestGetHTTPInternalServerError tests HTTP internal server error response.
func TestGetHTTPInternalServerError(t *testing.T) {
	actual := HTTPResponse{500, "Database Error", nil}
	expected := GetHTTPInternalServerError("Database Error")
	assert.Equal(t, expected, actual)

	actual = HTTPResponse{500, "Internal Server Error", nil}
	expected = GetHTTPInternalServerError("")
	assert.Equal(t, expected, actual)
}
