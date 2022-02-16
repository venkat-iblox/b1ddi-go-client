package convert_rname

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertRnameConvertRNameOK_Error(t *testing.T) {
	resp := NewConvertRnameConvertRNameOK()
	assert.NotEmpty(t, resp.Error())
}

func TestConvertRnameConvertRNameOK_GetPayload(t *testing.T) {
	resp := NewConvertRnameConvertRNameOK()
	assert.Nil(t, resp.GetPayload())
}
