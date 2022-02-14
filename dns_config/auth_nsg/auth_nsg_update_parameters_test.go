package auth_nsg

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

func TestNewAuthNsgUpdateParams(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthNsgUpdateParamsWithTimeout(t *testing.T) {
	p := NewAuthNsgUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthNsgUpdateParamsWithContext(t *testing.T) {
	p := NewAuthNsgUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthNsgUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthNsgUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgUpdateParams_WithDefaults(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	p = p.WithDefaults()
}

func TestAuthNsgUpdateParams_WithTimeout(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthNsgUpdateParams_WithContext(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthNsgUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgUpdateParams_WithBody(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	b := &models.ConfigAuthNSG{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestAuthNsgUpdateParams_WithID(t *testing.T) {
	p := NewAuthNsgUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
