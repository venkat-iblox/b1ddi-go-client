package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewListOK_Error(t *testing.T) {
	resp := NewViewListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestViewListOK_GetPayload(t *testing.T) {
	resp := NewViewListOK()
	assert.Nil(t, resp.GetPayload())
}
