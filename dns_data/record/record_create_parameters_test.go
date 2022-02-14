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

func TestNewRecordCreateParams(t *testing.T) {
	p := NewRecordCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordCreateParamsWithTimeout(t *testing.T) {
	p := NewRecordCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordCreateParamsWithContext(t *testing.T) {
	p := NewRecordCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordCreateParams_WithDefaults(t *testing.T) {
	p := NewRecordCreateParams()
	p = p.WithDefaults()
}

func TestRecordCreateParams_WithTimeout(t *testing.T) {
	p := NewRecordCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordCreateParams_WithContext(t *testing.T) {
	p := NewRecordCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordCreateParams_WithHTTPClient(t *testing.T) {
	p := NewRecordCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordCreateParams_WithBody(t *testing.T) {
	p := NewRecordCreateParams()
	b := &models.DataRecord{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
