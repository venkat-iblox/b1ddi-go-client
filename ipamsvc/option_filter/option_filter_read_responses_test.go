package option_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFilterReadOK_Error(t *testing.T) {
	resp := NewOptionFilterReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionFilterReadOK_GetPayload(t *testing.T) {
	resp := NewOptionFilterReadOK()
	assert.Nil(t, resp.GetPayload())
}
