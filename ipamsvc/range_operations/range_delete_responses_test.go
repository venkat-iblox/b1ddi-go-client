package range_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeDeleteNoContent_Error(t *testing.T) {
	resp := NewRangeDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
