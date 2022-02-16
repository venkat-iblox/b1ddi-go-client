package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceDeleteNoContent_Error(t *testing.T) {
	resp := NewIPSpaceDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
