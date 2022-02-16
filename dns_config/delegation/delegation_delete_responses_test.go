package delegation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelegationDeleteNoContent_Error(t *testing.T) {
	resp := NewDelegationDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
