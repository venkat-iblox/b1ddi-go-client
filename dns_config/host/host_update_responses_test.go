package host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostUpdateCreated_Error(t *testing.T) {
	resp := NewHostUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestHostUpdateCreated_GetPayload(t *testing.T) {
	resp := NewHostUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
