package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordCreateCreated_Error(t *testing.T) {
	resp := NewRecordCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestRecordCreateCreated_GetPayload(t *testing.T) {
	resp := NewRecordCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
