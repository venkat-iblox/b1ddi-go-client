package option_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionGroupReadOK_Error(t *testing.T) {
	resp := NewOptionGroupReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionGroupReadOK_GetPayload(t *testing.T) {
	resp := NewOptionGroupReadOK()
	assert.Nil(t, resp.GetPayload())
}
