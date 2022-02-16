package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeReadOK_Error(t *testing.T) {
	resp := NewRangeReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeReadOK_GetPayload(t *testing.T) {
	resp := NewRangeReadOK()
	assert.Nil(t, resp.GetPayload())
}
