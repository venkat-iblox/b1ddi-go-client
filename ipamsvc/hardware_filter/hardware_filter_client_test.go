package hardware_filter

import (
	"context"
	"crypto/tls"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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
		expectedRequest  http.Request
		testMethodName   string
		testMethodParams interface{}
	}{
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/hardware_filter"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Comment\"}\n")),
			},
			"HardwareFilterCreate",
			&HardwareFilterCreateParams{
				Body: &models.IpamsvcHardwareFilter{
					Comment: "Comment",
				},
				Context: context.TODO(),
			},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/hardware_filter/hardware-filter-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"HardwareFilterDelete",
			&HardwareFilterDeleteParams{
				ID:      "hardware-filter-delete-id",
				Context: context.TODO(),
			},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/hardware_filter"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"HardwareFilterList",
			&HardwareFilterListParams{
				Context: context.TODO(),
			},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/hardware_filter/hardware-filter-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"HardwareFilterRead",
			&HardwareFilterReadParams{
				ID:      "hardware-filter-read-id",
				Context: context.TODO(),
			},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/hardware_filter/hardware-filter-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			"HardwareFilterUpdate",
			&HardwareFilterUpdateParams{
				ID: "hardware-filter-update-id",
				Body: &models.IpamsvcHardwareFilter{
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
			c := initHardwareFilterTestClient(s.URL)

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

func initHardwareFilterTestClient(url string) ClientService {
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
