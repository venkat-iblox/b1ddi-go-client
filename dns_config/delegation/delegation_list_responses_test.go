package delegation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelegationListOK_Error(t *testing.T) {
	resp := NewDelegationListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDelegationListOK_GetPayload(t *testing.T) {
	resp := NewDelegationListOK()
	assert.Nil(t, resp.GetPayload())
}
