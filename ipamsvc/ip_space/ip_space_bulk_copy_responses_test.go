package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceBulkCopyCreated_Error(t *testing.T) {
	resp := NewIPSpaceBulkCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceBulkCopyCreated_GetPayload(t *testing.T) {
	resp := NewIPSpaceBulkCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
