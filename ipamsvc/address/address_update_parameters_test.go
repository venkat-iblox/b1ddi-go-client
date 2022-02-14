package address

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

func TestNewAddressUpdateParams(t *testing.T) {
	p := NewAddressUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressUpdateParamsWithTimeout(t *testing.T) {
	p := NewAddressUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressUpdateParamsWithContext(t *testing.T) {
	p := NewAddressUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressUpdateParams_WithDefaults(t *testing.T) {
	p := NewAddressUpdateParams()
	p = p.WithDefaults()
}

func TestAddressUpdateParams_WithTimeout(t *testing.T) {
	p := NewAddressUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressUpdateParams_WithContext(t *testing.T) {
	p := NewAddressUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewAddressUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressUpdateParams_WithBody(t *testing.T) {
	p := NewAddressUpdateParams()
	b := &models.IpamsvcAddress{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestAddressUpdateParams_WithID(t *testing.T) {
	p := NewAddressUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
