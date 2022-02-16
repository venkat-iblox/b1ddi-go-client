package dns_usage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDNSUsageReadOK_Error(t *testing.T) {
	resp := NewDNSUsageReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDNSUsageReadOK_GetPayload(t *testing.T) {
	resp := NewDNSUsageReadOK()
	assert.Nil(t, resp.GetPayload())
}
