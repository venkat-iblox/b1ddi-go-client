package ip_space

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

func TestNewIPSpaceBulkCopyParams(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceBulkCopyParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceBulkCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceBulkCopyParamsWithContext(t *testing.T) {
	p := NewIPSpaceBulkCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceBulkCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceBulkCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceBulkCopyParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	p = p.WithDefaults()
}

func TestIPSpaceBulkCopyParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceBulkCopyParams_WithContext(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceBulkCopyParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceBulkCopyParams_WithBody(t *testing.T) {
	p := NewIPSpaceBulkCopyParams()
	b := &models.IpamsvcBulkCopyIPSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
