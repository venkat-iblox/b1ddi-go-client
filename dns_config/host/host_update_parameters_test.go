package host

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

func TestNewHostUpdateParams(t *testing.T) {
	p := NewHostUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHostUpdateParamsWithTimeout(t *testing.T) {
	p := NewHostUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHostUpdateParamsWithContext(t *testing.T) {
	p := NewHostUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHostUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHostUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHostUpdateParams_WithDefaults(t *testing.T) {
	p := NewHostUpdateParams()
	p = p.WithDefaults()
}

func TestHostUpdateParams_WithTimeout(t *testing.T) {
	p := NewHostUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHostUpdateParams_WithContext(t *testing.T) {
	p := NewHostUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHostUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewHostUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHostUpdateParams_WithBody(t *testing.T) {
	p := NewHostUpdateParams()
	b := &models.ConfigHost{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestHostUpdateParams_WithID(t *testing.T) {
	p := NewHostUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
