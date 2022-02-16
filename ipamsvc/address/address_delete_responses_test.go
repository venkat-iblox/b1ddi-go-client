package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressDeleteNoContent_Error(t *testing.T) {
	resp := NewAddressDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
