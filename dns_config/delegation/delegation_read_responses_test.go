package delegation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelegationReadOK_Error(t *testing.T) {
	resp := NewDelegationReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDelegationReadOK_GetPayload(t *testing.T) {
	resp := NewDelegationReadOK()
	assert.Nil(t, resp.GetPayload())
}
