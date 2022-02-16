package option_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionCodeListOK_Error(t *testing.T) {
	resp := NewOptionCodeListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionCodeListOK_GetPayload(t *testing.T) {
	resp := NewOptionCodeListOK()
	assert.Nil(t, resp.GetPayload())
}
