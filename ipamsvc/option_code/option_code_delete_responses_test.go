package option_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionCodeDeleteNoContent_Error(t *testing.T) {
	resp := NewOptionCodeDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
