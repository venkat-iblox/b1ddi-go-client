package option_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionGroupDeleteNoContent_Error(t *testing.T) {
	resp := NewOptionGroupDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
