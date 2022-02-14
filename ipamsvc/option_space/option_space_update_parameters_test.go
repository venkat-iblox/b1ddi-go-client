package option_space

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

func TestNewOptionSpaceUpdateParams(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionSpaceUpdateParamsWithTimeout(t *testing.T) {
	p := NewOptionSpaceUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionSpaceUpdateParamsWithContext(t *testing.T) {
	p := NewOptionSpaceUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionSpaceUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionSpaceUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceUpdateParams_WithDefaults(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	p = p.WithDefaults()
}

func TestOptionSpaceUpdateParams_WithTimeout(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionSpaceUpdateParams_WithContext(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionSpaceUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceUpdateParams_WithBody(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	b := &models.IpamsvcOptionSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestOptionSpaceUpdateParams_WithID(t *testing.T) {
	p := NewOptionSpaceUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
