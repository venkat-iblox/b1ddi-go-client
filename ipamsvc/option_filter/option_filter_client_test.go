package option_filter

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
		expectedRequest  http.Request
		testMethodParams interface{}
	}{
		{
			"OptionFilterCreate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_filter"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name\":\"test_name\"}\n")),
			},
			&OptionFilterCreateParams{
				Body:    &models.IpamsvcOptionFilter{Name: swag.String("test_name")},
				Context: context.TODO(),
			},
		},
		{
			"OptionFilterDelete",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_filter/option-filter-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionFilterDeleteParams{
				ID:      "option-filter-delete-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionFilterList",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_filter"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionFilterListParams{
				Context: context.TODO(),
			},
		},
		{
			"OptionFilterRead",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_filter/option-filter-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionFilterReadParams{
				ID:      "option-filter-read-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionFilterUpdate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_filter/option-filter-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			&OptionFilterUpdateParams{
				ID: "option-filter-update-id",
				Body: &models.IpamsvcOptionFilter{
					Comment: "Updated comment",
				},
				Context: context.TODO(),
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
			c := initOptionFilterTestClient(s.URL)

			// Compose test function call parameters
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

func initOptionFilterTestClient(url string) ClientService {
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
