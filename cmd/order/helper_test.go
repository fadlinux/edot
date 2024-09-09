package order

import (
	"testing"
)

func Test_newDBConnection(t *testing.T) {
	newDBConnection("mysql", "host")
}
