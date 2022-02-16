package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordUpdateCreated_Error(t *testing.T) {
	resp := NewRecordUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRecordUpdateCreated_GetPayload(t *testing.T) {
	resp := NewRecordUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
