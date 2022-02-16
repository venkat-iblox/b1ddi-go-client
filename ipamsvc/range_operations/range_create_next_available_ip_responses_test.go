package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeCreateNextAvailableIPCreated_Error(t *testing.T) {
	resp := NewRangeCreateNextAvailableIPCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeCreateNextAvailableIPCreated_GetPayload(t *testing.T) {
	resp := NewRangeCreateNextAvailableIPCreated()
	assert.Nil(t, resp.GetPayload())
}
