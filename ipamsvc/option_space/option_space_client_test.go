package option_space

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
			"OptionSpaceCreate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_space"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name\":\"test_name\"}\n")),
			},
			&OptionSpaceCreateParams{
				Body: &models.IpamsvcOptionSpace{
					Name: swag.String("test_name"),
				},
				Context: context.TODO(),
			},
		},
		{
			"OptionSpaceDelete",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_space/option-space-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionSpaceDeleteParams{
				ID:      "option-space-delete-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionSpaceList",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_space"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionSpaceListParams{
				Context: context.TODO(),
			},
		},
		{
			"OptionSpaceRead",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_space/option-space-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionSpaceReadParams{
				ID:      "option-space-read-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionSpaceUpdate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_space/option-space-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			&OptionSpaceUpdateParams{
				ID: "option-space-update-id",
				Body: &models.IpamsvcOptionSpace{
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
			c := initOptionSpaceTestClient(s.URL)

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

func initOptionSpaceTestClient(url string) ClientService {
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
