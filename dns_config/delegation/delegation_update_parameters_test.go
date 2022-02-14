package delegation

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

func TestNewDelegationUpdateParams(t *testing.T) {
	p := NewDelegationUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDelegationUpdateParamsWithTimeout(t *testing.T) {
	p := NewDelegationUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDelegationUpdateParamsWithContext(t *testing.T) {
	p := NewDelegationUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDelegationUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDelegationUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationUpdateParams_WithDefaults(t *testing.T) {
	p := NewDelegationUpdateParams()
	p = p.WithDefaults()
}

func TestDelegationUpdateParams_WithTimeout(t *testing.T) {
	p := NewDelegationUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDelegationUpdateParams_WithContext(t *testing.T) {
	p := NewDelegationUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDelegationUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewDelegationUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationUpdateParams_WithBody(t *testing.T) {
	p := NewDelegationUpdateParams()
	b := &models.ConfigDelegation{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestDelegationUpdateParams_WithID(t *testing.T) {
	p := NewDelegationUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
