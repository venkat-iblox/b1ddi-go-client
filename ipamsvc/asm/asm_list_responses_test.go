package asm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsmListOK_Error(t *testing.T) {
	resp := NewAsmListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAsmListOK_GetPayload(t *testing.T) {
	resp := NewAsmListOK()
	assert.Nil(t, resp.GetPayload())
}
