package acl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACLUpdateCreated_Error(t *testing.T) {
	resp := NewACLUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestACLUpdateCreated_GetPayload(t *testing.T) {
	resp := NewACLUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
