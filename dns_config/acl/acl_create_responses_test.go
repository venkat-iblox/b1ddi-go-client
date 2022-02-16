package acl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACLCreateCreated_Error(t *testing.T) {
	resp := NewACLCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestACLCreateCreated_GetPayload(t *testing.T) {
	resp := NewACLCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
