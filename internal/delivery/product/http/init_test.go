package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductHTTP(t *testing.T) {
	delivery := NewProductHTTP(nil)
	assert.NotNil(t, delivery)
}
