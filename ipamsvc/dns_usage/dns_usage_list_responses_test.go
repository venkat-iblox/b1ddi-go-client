package dns_usage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDNSUsageListOK_Error(t *testing.T) {
	resp := NewDNSUsageListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDNSUsageListOK_GetPayload(t *testing.T) {
	resp := NewDNSUsageListOK()
	assert.Nil(t, resp.GetPayload())
}
