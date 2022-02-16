package asm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsmCreateCreated_Error(t *testing.T) {
	resp := NewAsmCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAsmCreateCreated_GetPayload(t *testing.T) {
	resp := NewAsmCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
