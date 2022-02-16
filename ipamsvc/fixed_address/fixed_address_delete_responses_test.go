package fixed_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedAddressDeleteNoContent_Error(t *testing.T) {
	resp := NewFixedAddressDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
