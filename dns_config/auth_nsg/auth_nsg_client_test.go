package auth_nsg

import (
	"context"
	"crypto/tls"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/infobloxopen/b1ddi-go-client/runtimetest"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		testMethodName   string
		testMethodParams interface{}
		expectedRequest  http.Request
	}{
		{
			"AuthNsgCreate",
			&AuthNsgCreateParams{
				Body: &models.ConfigAuthNSG{
					Name: swag.String("test_auth_nsg"),
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_nsg"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name\":\"test_auth_nsg\"}\n")),
			},
		},
		{
			"AuthNsgDelete",
			&AuthNsgDeleteParams{
				ID:      "auth-nsg-delete-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_nsg/auth-nsg-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthNsgList",
			&AuthNsgListParams{
				Fields:    swag.String("field"),
				Filter:    swag.String("filter"),
				Limit:     swag.Int64(int64(20)),
				Offset:    swag.Int64(int64(20)),
				OrderBy:   swag.String("desc"),
				PageToken: swag.String("token"),
				Tfilter:   swag.String("tfilter"),
				TorderBy:  swag.String("desc"),
				Context:   context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/dns/auth_nsg",
					RawQuery: "_fields=field&_filter=filter&_limit=20&_offset=20&_order_by=desc&_page_token=token&_tfilter=tfilter&_torder_by=desc",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthNsgRead",
			&AuthNsgReadParams{
				ID:      "auth-nsg-read-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_nsg/auth-nsg-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthNsgUpdate",
			&AuthNsgUpdateParams{
				ID: "auth-nsg-update-id",
				Body: &models.ConfigAuthNSG{
					Name:    swag.String("test_auth_nsg"),
					Comment: "Updated comment",
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_nsg/auth-nsg-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\",\"name\":\"test_auth_nsg\"}\n")),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testMethodName, func(t *testing.T) {
			// Get the mock server response body
			resp := reflect.ValueOf(&Client{}).MethodByName(tc.testMethodName).Type().Out(0)
			// Initialize mock server
			s := runtimetest.InitTestServer(t, tc.expectedRequest, reflect.New(resp))

			// Initialize the client
			c := initAuthNSGTestClient(s.URL)

			methodParams := []reflect.Value{
				reflect.ValueOf(tc.testMethodParams),
				reflect.New(reflect.TypeOf((*runtime.ClientAuthInfoWriter)(nil)).Elem()).Elem(),
			}

			// Call the method by the name specified in the test case
			results := reflect.ValueOf(c).MethodByName(tc.testMethodName).Call(methodParams)
			if err := results[1].Interface(); err != nil {
				t.Fatal(err)
			}

			// Close the mock server
			s.Close()
		})
	}
}

func initAuthNSGTestClient(url string) ClientService {
	transport := httptransport.New(
		strings.TrimPrefix(url, "https://"), "api/ddi/v1", nil,
	)
	// Disable TLS verify for the mock server
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	transport.Transport = tr
	// create the Address API client
	return New(transport, strfmt.Default)
}
