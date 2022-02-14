package convert_rname

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewConvertRnameConvertRNameParams(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewConvertRnameConvertRNameParamsWithTimeout(t *testing.T) {
	p := NewConvertRnameConvertRNameParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewConvertRnameConvertRNameParamsWithContext(t *testing.T) {
	p := NewConvertRnameConvertRNameParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewConvertRnameConvertRNameParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewConvertRnameConvertRNameParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestConvertRnameConvertRNameParams_WithDefaults(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	p = p.WithDefaults()
}

func TestConvertRnameConvertRNameParams_WithTimeout(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestConvertRnameConvertRNameParams_WithContext(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestConvertRnameConvertRNameParams_WithHTTPClient(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestConvertRnameConvertRNameParams_WithEmailAddress(t *testing.T) {
	p := NewConvertRnameConvertRNameParams()
	p = p.WithEmailAddress("email")
	assert.Equal(t, "email", p.EmailAddress)
}
