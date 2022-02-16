package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetDeleteNoContent_Error(t *testing.T) {
	resp := NewSubnetDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
