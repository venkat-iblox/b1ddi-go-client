package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordListOK_Error(t *testing.T) {
	resp := NewRecordListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestRecordListOK_GetPayload(t *testing.T) {
	resp := NewRecordListOK()
	assert.Nil(t, resp.GetPayload())
}
