package ha_group

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

func TestNewHaGroupCreateParams(t *testing.T) {
	p := NewHaGroupCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHaGroupCreateParamsWithTimeout(t *testing.T) {
	p := NewHaGroupCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHaGroupCreateParamsWithContext(t *testing.T) {
	p := NewHaGroupCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHaGroupCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHaGroupCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupCreateParams_WithDefaults(t *testing.T) {
	p := NewHaGroupCreateParams()
	p = p.WithDefaults()
}

func TestHaGroupCreateParams_WithTimeout(t *testing.T) {
	p := NewHaGroupCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHaGroupCreateParams_WithContext(t *testing.T) {
	p := NewHaGroupCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHaGroupCreateParams_WithHTTPClient(t *testing.T) {
	p := NewHaGroupCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupCreateParams_WithBody(t *testing.T) {
	p := NewHaGroupCreateParams()
	b := &models.IpamsvcHAGroup{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
