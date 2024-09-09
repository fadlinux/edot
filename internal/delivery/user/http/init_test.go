package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserHTTP(t *testing.T) {
	delivery := NewUserHTTP(nil)
	assert.NotNil(t, delivery)
}
