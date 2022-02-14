package ha_group

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

func TestNewHaGroupReadParams(t *testing.T) {
	p := NewHaGroupReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHaGroupReadParamsWithTimeout(t *testing.T) {
	p := NewHaGroupReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHaGroupReadParamsWithContext(t *testing.T) {
	p := NewHaGroupReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHaGroupReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHaGroupReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupReadParams_WithDefaults(t *testing.T) {
	p := NewHaGroupReadParams()
	p = p.WithDefaults()
}

func TestHaGroupReadParams_WithTimeout(t *testing.T) {
	p := NewHaGroupReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHaGroupReadParams_WithContext(t *testing.T) {
	p := NewHaGroupReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHaGroupReadParams_WithHTTPClient(t *testing.T) {
	p := NewHaGroupReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupReadParams_WithFields(t *testing.T) {
	p := NewHaGroupReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestHaGroupReadParams_WithID(t *testing.T) {
	p := NewHaGroupReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
