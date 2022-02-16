package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeCreateCreated_Error(t *testing.T) {
	resp := NewRangeCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeCreateCreated_GetPayload(t *testing.T) {
	resp := NewRangeCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
