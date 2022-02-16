package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeUpdateCreated_Error(t *testing.T) {
	resp := NewRangeUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeUpdateCreated_GetPayload(t *testing.T) {
	resp := NewRangeUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
