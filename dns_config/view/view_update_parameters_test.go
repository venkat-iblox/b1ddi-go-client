package view

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

func TestNewViewUpdateParams(t *testing.T) {
	p := NewViewUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewViewUpdateParamsWithTimeout(t *testing.T) {
	p := NewViewUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewViewUpdateParamsWithContext(t *testing.T) {
	p := NewViewUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewViewUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewViewUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewUpdateParams_WithDefaults(t *testing.T) {
	p := NewViewUpdateParams()
	p = p.WithDefaults()
}

func TestViewUpdateParams_WithTimeout(t *testing.T) {
	p := NewViewUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestViewUpdateParams_WithContext(t *testing.T) {
	p := NewViewUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestViewUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewViewUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewUpdateParams_WithBody(t *testing.T) {
	p := NewViewUpdateParams()
	b := &models.ConfigView{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestViewUpdateParams_WithID(t *testing.T) {
	p := NewViewUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
