package acl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACLReadOK_Error(t *testing.T) {
	resp := NewACLReadOK()
	assert.NotEmpty(t, resp.Error())
}
