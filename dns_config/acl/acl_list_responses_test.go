package acl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACLListOK_Error(t *testing.T) {
	resp := NewACLListOK()
	assert.NotEmpty(t, resp.Error())
}
