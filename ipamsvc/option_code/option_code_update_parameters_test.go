package option_code

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

func TestNewOptionCodeUpdateParams(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionCodeUpdateParamsWithTimeout(t *testing.T) {
	p := NewOptionCodeUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionCodeUpdateParamsWithContext(t *testing.T) {
	p := NewOptionCodeUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionCodeUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionCodeUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeUpdateParams_WithDefaults(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	p = p.WithDefaults()
}

func TestOptionCodeUpdateParams_WithTimeout(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionCodeUpdateParams_WithContext(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionCodeUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeUpdateParams_WithBody(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	b := &models.IpamsvcOptionCode{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestOptionCodeUpdateParams_WithID(t *testing.T) {
	p := NewOptionCodeUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
