package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordSOASerialIncrementCreated_Error(t *testing.T) {
	resp := NewRecordSOASerialIncrementCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRecordSOASerialIncrementCreated_GetPayload(t *testing.T) {
	resp := NewRecordSOASerialIncrementCreated()
	assert.Nil(t, resp.GetPayload())
}
