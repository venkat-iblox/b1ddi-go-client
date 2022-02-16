package ha_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaGroupDeleteNoContent_Error(t *testing.T) {
	resp := NewHaGroupDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
