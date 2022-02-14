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

func TestNewIPSpaceUpdateParams(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceUpdateParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceUpdateParamsWithContext(t *testing.T) {
	p := NewIPSpaceUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceUpdateParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	p = p.WithDefaults()
}

func TestIPSpaceUpdateParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceUpdateParams_WithContext(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceUpdateParams_WithBody(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	b := &models.IpamsvcIPSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestIPSpaceUpdateParams_WithID(t *testing.T) {
	p := NewIPSpaceUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
