package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordReadOK_Error(t *testing.T) {
	resp := NewRecordReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestRecordReadOK_GetPayload(t *testing.T) {
	resp := NewRecordReadOK()
	assert.Nil(t, resp.GetPayload())
}
