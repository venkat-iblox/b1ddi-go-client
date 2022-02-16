package host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostListOK_Error(t *testing.T) {
	resp := NewHostListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHostListOK_GetPayload(t *testing.T) {
	resp := NewHostListOK()
	assert.Nil(t, resp.GetPayload())
}
