package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShopHTTP(t *testing.T) {
	delivery := NewShopHTTP(nil)
	assert.NotNil(t, delivery)
}
