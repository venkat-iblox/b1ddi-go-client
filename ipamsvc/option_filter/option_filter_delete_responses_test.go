package option_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFilterDeleteNoContent_Error(t *testing.T) {
	resp := NewOptionFilterDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
