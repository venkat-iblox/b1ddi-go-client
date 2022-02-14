package record

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

func TestNewRecordUpdateParams(t *testing.T) {
	p := NewRecordUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordUpdateParamsWithTimeout(t *testing.T) {
	p := NewRecordUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordUpdateParamsWithContext(t *testing.T) {
	p := NewRecordUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordUpdateParams_WithDefaults(t *testing.T) {
	p := NewRecordUpdateParams()
	p = p.WithDefaults()
}

func TestRecordUpdateParams_WithTimeout(t *testing.T) {
	p := NewRecordUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordUpdateParams_WithContext(t *testing.T) {
	p := NewRecordUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewRecordUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordUpdateParams_WithBody(t *testing.T) {
	p := NewRecordUpdateParams()
	b := &models.DataRecord{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestRecordUpdateParams_WithID(t *testing.T) {
	p := NewRecordUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
