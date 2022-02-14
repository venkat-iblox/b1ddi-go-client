package delegation

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

func TestNewDelegationCreateParams(t *testing.T) {
	p := NewDelegationCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDelegationCreateParamsWithTimeout(t *testing.T) {
	p := NewDelegationCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDelegationCreateParamsWithContext(t *testing.T) {
	p := NewDelegationCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDelegationCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDelegationCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationCreateParams_WithDefaults(t *testing.T) {
	p := NewDelegationCreateParams()
	p = p.WithDefaults()
}

func TestDelegationCreateParams_WithTimeout(t *testing.T) {
	p := NewDelegationCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDelegationCreateParams_WithContext(t *testing.T) {
	p := NewDelegationCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDelegationCreateParams_WithHTTPClient(t *testing.T) {
	p := NewDelegationCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationCreateParams_WithBody(t *testing.T) {
	p := NewDelegationCreateParams()
	b := &models.ConfigDelegation{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
