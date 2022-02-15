package view

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

func TestNewViewBulkCopyParams(t *testing.T) {
	p := NewViewBulkCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewViewBulkCopyParamsWithTimeout(t *testing.T) {
	p := NewViewBulkCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewViewBulkCopyParamsWithContext(t *testing.T) {
	p := NewViewBulkCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewViewBulkCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewViewBulkCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewBulkCopyParams_WithDefaults(t *testing.T) {
	p := NewViewBulkCopyParams()
	p = p.WithDefaults()
}

func TestViewBulkCopyParams_WithTimeout(t *testing.T) {
	p := NewViewBulkCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestViewBulkCopyParams_WithContext(t *testing.T) {
	p := NewViewBulkCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestViewBulkCopyParams_WithHTTPClient(t *testing.T) {
	p := NewViewBulkCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewBulkCopyParams_WithBody(t *testing.T) {
	p := NewViewBulkCopyParams()
	b := &models.ConfigBulkCopyView{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
