package delegation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelegationUpdateCreated_Error(t *testing.T) {
	resp := NewDelegationUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestDelegationUpdateCreated_GetPayload(t *testing.T) {
	resp := NewDelegationUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
