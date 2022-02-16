package host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostReadOK_Error(t *testing.T) {
	resp := NewHostReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHostReadOK_GetPayload(t *testing.T) {
	resp := NewHostReadOK()
	assert.Nil(t, resp.GetPayload())
}
