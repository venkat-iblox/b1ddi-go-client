package range_operations

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

func TestNewRangeUpdateParams(t *testing.T) {
	p := NewRangeUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeUpdateParamsWithTimeout(t *testing.T) {
	p := NewRangeUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeUpdateParamsWithContext(t *testing.T) {
	p := NewRangeUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeUpdateParams_WithDefaults(t *testing.T) {
	p := NewRangeUpdateParams()
	p = p.WithDefaults()
}

func TestRangeUpdateParams_WithTimeout(t *testing.T) {
	p := NewRangeUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeUpdateParams_WithContext(t *testing.T) {
	p := NewRangeUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewRangeUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeUpdateParams_WithBody(t *testing.T) {
	p := NewRangeUpdateParams()
	b := &models.IpamsvcRange{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestRangeUpdateParams_WithID(t *testing.T) {
	p := NewRangeUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
