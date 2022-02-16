package asm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsmReadOK_Error(t *testing.T) {
	resp := NewAsmReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAsmReadOK_GetPayload(t *testing.T) {
	resp := NewAsmReadOK()
	assert.Nil(t, resp.GetPayload())
}
