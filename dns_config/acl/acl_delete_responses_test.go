package acl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACLDeleteNoContent_Error(t *testing.T) {
	resp := NewACLDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
