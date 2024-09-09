package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderHTTP(t *testing.T) {
	delivery := NewOrderHTTP(nil)
	assert.NotNil(t, delivery)
}
