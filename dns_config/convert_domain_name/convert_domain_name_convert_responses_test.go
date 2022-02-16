package convert_domain_name

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertDomainNameConvertOK_Error(t *testing.T) {
	resp := NewConvertDomainNameConvertOK()
	assert.NotEmpty(t, resp.Error())
}

func TestConvertDomainNameConvertOK_GetPayload(t *testing.T) {
	resp := NewConvertDomainNameConvertOK()
	assert.Nil(t, resp.GetPayload())
}
