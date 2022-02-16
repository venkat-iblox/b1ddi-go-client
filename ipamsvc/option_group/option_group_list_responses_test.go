package option_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionGroupListOK_Error(t *testing.T) {
	resp := NewOptionGroupListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionGroupListOK_GetPayload(t *testing.T) {
	resp := NewOptionGroupListOK()
	assert.Nil(t, resp.GetPayload())
}
