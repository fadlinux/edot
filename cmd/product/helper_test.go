package product

import (
	"testing"
)

func Test_newDBConnection(t *testing.T) {
	newDBConnection("mysql", "host")
}
