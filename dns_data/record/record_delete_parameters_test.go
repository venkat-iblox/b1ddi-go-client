package record

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewRecordDeleteParams(t *testing.T) {
	p := NewRecordDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordDeleteParamsWithTimeout(t *testing.T) {
	p := NewRecordDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordDeleteParamsWithContext(t *testing.T) {
	p := NewRecordDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordDeleteParams_WithDefaults(t *testing.T) {
	p := NewRecordDeleteParams()
	p = p.WithDefaults()
}

func TestRecordDeleteParams_WithTimeout(t *testing.T) {
	p := NewRecordDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordDeleteParams_WithContext(t *testing.T) {
	p := NewRecordDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewRecordDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordDeleteParams_WithID(t *testing.T) {
	p := NewRecordDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
