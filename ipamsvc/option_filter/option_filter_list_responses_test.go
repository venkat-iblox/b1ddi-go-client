package option_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFilterListOK_Error(t *testing.T) {
	resp := NewOptionFilterListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionFilterListOK_GetPayload(t *testing.T) {
	resp := NewOptionFilterListOK()
	assert.Nil(t, resp.GetPayload())
}
