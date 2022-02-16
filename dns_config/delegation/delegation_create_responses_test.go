package delegation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelegationCreateCreated_Error(t *testing.T) {
	resp := NewDelegationCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestDelegationCreateCreated_GetPayload(t *testing.T) {
	resp := NewDelegationCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
