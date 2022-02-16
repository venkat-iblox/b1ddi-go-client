package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeListNextAvailableIPOK_Error(t *testing.T) {
	resp := NewRangeListNextAvailableIPOK()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeListNextAvailableIPOK_GetPayload(t *testing.T) {
	resp := NewRangeListNextAvailableIPOK()
	assert.Nil(t, resp.GetPayload())
}
