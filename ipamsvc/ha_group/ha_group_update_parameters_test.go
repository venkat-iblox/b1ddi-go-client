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

func TestNewHaGroupUpdateParams(t *testing.T) {
	p := NewHaGroupUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHaGroupUpdateParamsWithTimeout(t *testing.T) {
	p := NewHaGroupUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHaGroupUpdateParamsWithContext(t *testing.T) {
	p := NewHaGroupUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHaGroupUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHaGroupUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupUpdateParams_WithDefaults(t *testing.T) {
	p := NewHaGroupUpdateParams()
	p = p.WithDefaults()
}

func TestHaGroupUpdateParams_WithTimeout(t *testing.T) {
	p := NewHaGroupUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHaGroupUpdateParams_WithContext(t *testing.T) {
	p := NewHaGroupUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHaGroupUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewHaGroupUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupUpdateParams_WithBody(t *testing.T) {
	p := NewHaGroupUpdateParams()
	b := &models.IpamsvcHAGroup{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestHaGroupUpdateParams_WithID(t *testing.T) {
	p := NewHaGroupUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
