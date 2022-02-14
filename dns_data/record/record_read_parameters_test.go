package record

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewRecordReadParams(t *testing.T) {
	p := NewRecordReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordReadParamsWithTimeout(t *testing.T) {
	p := NewRecordReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordReadParamsWithContext(t *testing.T) {
	p := NewRecordReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordReadParams_WithDefaults(t *testing.T) {
	p := NewRecordReadParams()
	p = p.WithDefaults()
}

func TestRecordReadParams_WithTimeout(t *testing.T) {
	p := NewRecordReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordReadParams_WithContext(t *testing.T) {
	p := NewRecordReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordReadParams_WithHTTPClient(t *testing.T) {
	p := NewRecordReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordReadParams_WithFields(t *testing.T) {
	p := NewRecordReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestRecordReadParams_WithID(t *testing.T) {
	p := NewRecordReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
