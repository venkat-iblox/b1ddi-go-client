package forward_nsg

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

func TestNewForwardNsgUpdateParams(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardNsgUpdateParamsWithTimeout(t *testing.T) {
	p := NewForwardNsgUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardNsgUpdateParamsWithContext(t *testing.T) {
	p := NewForwardNsgUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardNsgUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardNsgUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgUpdateParams_WithDefaults(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	p = p.WithDefaults()
}

func TestForwardNsgUpdateParams_WithTimeout(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardNsgUpdateParams_WithContext(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardNsgUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgUpdateParams_WithBody(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	b := &models.ConfigForwardNSG{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestForwardNsgUpdateParams_WithID(t *testing.T) {
	p := NewForwardNsgUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
