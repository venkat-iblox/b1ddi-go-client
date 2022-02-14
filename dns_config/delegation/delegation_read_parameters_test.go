package delegation

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

func TestNewDelegationReadParams(t *testing.T) {
	p := NewDelegationReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDelegationReadParamsWithTimeout(t *testing.T) {
	p := NewDelegationReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDelegationReadParamsWithContext(t *testing.T) {
	p := NewDelegationReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDelegationReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDelegationReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationReadParams_WithDefaults(t *testing.T) {
	p := NewDelegationReadParams()
	p = p.WithDefaults()
}

func TestDelegationReadParams_WithTimeout(t *testing.T) {
	p := NewDelegationReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDelegationReadParams_WithContext(t *testing.T) {
	p := NewDelegationReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDelegationReadParams_WithHTTPClient(t *testing.T) {
	p := NewDelegationReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationReadParams_WithFields(t *testing.T) {
	p := NewDelegationReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDelegationReadParams_WithID(t *testing.T) {
	p := NewDelegationReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
