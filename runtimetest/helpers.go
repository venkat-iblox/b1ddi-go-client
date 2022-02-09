package runtimetest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitTestServer(t *testing.T, expectedRequest http.Request, respBody interface{}) *httptest.Server {
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, expectedRequest.URL.Path, r.URL.Path)
		assert.Equal(t, expectedRequest.Method, r.Method)

		expectedBody, err := io.ReadAll(expectedRequest.Body)
		if err != nil {
			t.Error(err)
		}
		actualBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedBody, actualBody)

		assert.NotNil(t, respBody)

		w.Header().Set("content-type", "application/json")
		body, err := json.Marshal(respBody)
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write(body)
		if err != nil {
			t.Error(err)
		}
	}))
}
