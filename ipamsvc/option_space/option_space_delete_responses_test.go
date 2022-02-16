package option_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionSpaceDeleteNoContent_Error(t *testing.T) {
	resp := NewOptionSpaceDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
