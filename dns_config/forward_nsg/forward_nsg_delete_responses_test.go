package forward_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardNsgDeleteNoContent_Error(t *testing.T) {
	resp := NewForwardNsgDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
