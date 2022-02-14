package asm

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

func TestNewAsmListParams(t *testing.T) {
	p := NewAsmListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAsmListParamsWithTimeout(t *testing.T) {
	p := NewAsmListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAsmListParamsWithContext(t *testing.T) {
	p := NewAsmListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAsmListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAsmListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmListParams_WithDefaults(t *testing.T) {
	p := NewAsmListParams()
	p = p.WithDefaults()
}

func TestAsmListParams_WithTimeout(t *testing.T) {
	p := NewAsmListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAsmListParams_WithContext(t *testing.T) {
	p := NewAsmListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAsmListParams_WithHTTPClient(t *testing.T) {
	p := NewAsmListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmListParams_WithFields(t *testing.T) {
	p := NewAsmListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAsmListParams_WithSubnetID(t *testing.T) {
	p := NewAsmListParams()
	p = p.WithSubnetID(swag.String("subnet-id"))
	require.NotNil(t, p.SubnetID)
	assert.Equal(t, "subnet-id", *p.SubnetID)
}
