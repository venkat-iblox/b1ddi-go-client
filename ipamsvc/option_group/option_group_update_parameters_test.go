package option_group

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

func TestNewOptionGroupUpdateParams(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionGroupUpdateParamsWithTimeout(t *testing.T) {
	p := NewOptionGroupUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionGroupUpdateParamsWithContext(t *testing.T) {
	p := NewOptionGroupUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionGroupUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionGroupUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupUpdateParams_WithDefaults(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	p = p.WithDefaults()
}

func TestOptionGroupUpdateParams_WithTimeout(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionGroupUpdateParams_WithContext(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionGroupUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupUpdateParams_WithBody(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	b := &models.IpamsvcOptionGroup{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestOptionGroupUpdateParams_WithID(t *testing.T) {
	p := NewOptionGroupUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
