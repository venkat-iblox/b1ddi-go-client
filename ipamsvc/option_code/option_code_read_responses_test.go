package option_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionCodeReadOK_Error(t *testing.T) {
	resp := NewOptionCodeReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionCodeReadOK_GetPayload(t *testing.T) {
	resp := NewOptionCodeReadOK()
	assert.Nil(t, resp.GetPayload())
}
