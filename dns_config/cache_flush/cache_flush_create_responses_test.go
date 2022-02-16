package cache_flush

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheFlushCreateCreated_Error(t *testing.T) {
	resp := NewCacheFlushCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestCacheFlushCreateCreated_GetPayload(t *testing.T) {
	resp := NewCacheFlushCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
