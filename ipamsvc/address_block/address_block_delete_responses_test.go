package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockDeleteNoContent_Error(t *testing.T) {
	resp := NewAddressBlockDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
