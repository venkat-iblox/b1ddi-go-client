package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterListOK_Error(t *testing.T) {
	resp := NewFilterListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestFilterListOK_GetPayload(t *testing.T) {
	resp := NewFilterListOK()
	assert.Nil(t, resp.GetPayload())
}
