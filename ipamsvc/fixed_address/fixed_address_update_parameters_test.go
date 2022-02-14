package fixed_address

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

func TestNewFixedAddressUpdateParams(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFixedAddressUpdateParamsWithTimeout(t *testing.T) {
	p := NewFixedAddressUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFixedAddressUpdateParamsWithContext(t *testing.T) {
	p := NewFixedAddressUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFixedAddressUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFixedAddressUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressUpdateParams_WithDefaults(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	p = p.WithDefaults()
}

func TestFixedAddressUpdateParams_WithTimeout(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFixedAddressUpdateParams_WithContext(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFixedAddressUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressUpdateParams_WithBody(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	b := &models.IpamsvcFixedAddress{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestFixedAddressUpdateParams_WithID(t *testing.T) {
	p := NewFixedAddressUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
