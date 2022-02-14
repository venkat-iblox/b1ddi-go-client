package delegation

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewDelegationDeleteParams(t *testing.T) {
	p := NewDelegationDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDelegationDeleteParamsWithTimeout(t *testing.T) {
	p := NewDelegationDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDelegationDeleteParamsWithContext(t *testing.T) {
	p := NewDelegationDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDelegationDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDelegationDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationDeleteParams_WithDefaults(t *testing.T) {
	p := NewDelegationDeleteParams()
	p = p.WithDefaults()
}

func TestDelegationDeleteParams_WithTimeout(t *testing.T) {
	p := NewDelegationDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDelegationDeleteParams_WithContext(t *testing.T) {
	p := NewDelegationDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDelegationDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewDelegationDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationDeleteParams_WithID(t *testing.T) {
	p := NewDelegationDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
