package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordDeleteNoContent_Error(t *testing.T) {
	resp := NewRecordDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
