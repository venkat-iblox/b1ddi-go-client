package record

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewRecordSOASerialIncrementParams(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	require.NotNil(t, p)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordSOASerialIncrementParamsWithTimeout(t *testing.T) {
	p := NewRecordSOASerialIncrementParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordSOASerialIncrementParamsWithContext(t *testing.T) {
	p := NewRecordSOASerialIncrementParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordSOASerialIncrementParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordSOASerialIncrementParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordSOASerialIncrementParams_WithDefaults(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	p = p.WithDefaults()
}

func TestRecordSOASerialIncrementParams_WithTimeout(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordSOASerialIncrementParams_WithContext(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordSOASerialIncrementParams_WithHTTPClient(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordSOASerialIncrementParams_WithBody(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	b := &models.DataSOASerialIncrementRequest{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestRecordSOASerialIncrementParams_WithID(t *testing.T) {
	p := NewRecordSOASerialIncrementParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
