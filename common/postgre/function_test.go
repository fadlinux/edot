package postgre

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestNewDBConnection(t *testing.T) {
	NewDBConnection("postgres", "host")
}
