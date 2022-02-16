package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeListOK_Error(t *testing.T) {
	resp := NewRangeListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestRangeListOK_GetPayload(t *testing.T) {
	resp := NewRangeListOK()
	assert.Nil(t, resp.GetPayload())
}
