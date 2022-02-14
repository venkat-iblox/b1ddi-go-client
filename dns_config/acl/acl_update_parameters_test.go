package acl

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

func TestNewACLUpdateParams(t *testing.T) {
	p := NewACLUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewACLUpdateParamsWithTimeout(t *testing.T) {
	p := NewACLUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewACLUpdateParamsWithContext(t *testing.T) {
	p := NewACLUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewACLUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewACLUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLUpdateParams_WithDefaults(t *testing.T) {
	p := NewACLUpdateParams()
	p = p.WithDefaults()
}

func TestACLUpdateParams_WithTimeout(t *testing.T) {
	p := NewACLUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestACLUpdateParams_WithContext(t *testing.T) {
	p := NewACLUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestACLUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewACLUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLUpdateParams_WithBody(t *testing.T) {
	p := NewACLUpdateParams()
	b := &models.ConfigACL{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestACLUpdateParams_WithID(t *testing.T) {
	p := NewACLUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
