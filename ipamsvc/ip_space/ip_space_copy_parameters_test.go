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

func TestNewIPSpaceCopyParams(t *testing.T) {
	p := NewIPSpaceCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceCopyParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceCopyParamsWithContext(t *testing.T) {
	p := NewIPSpaceCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceCopyParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceCopyParams()
	p = p.WithDefaults()
}

func TestIPSpaceCopyParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceCopyParams_WithContext(t *testing.T) {
	p := NewIPSpaceCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceCopyParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceCopyParams_WithBody(t *testing.T) {
	p := NewIPSpaceCopyParams()
	b := &models.IpamsvcCopyIPSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestIPSpaceCopyParams_WithID(t *testing.T) {
	p := NewIPSpaceCopyParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
