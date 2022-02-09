package option_group

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
			"OptionGroupCreate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_group"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name\":\"test_name\"}\n")),
			},
			&OptionGroupCreateParams{
				Body: &models.IpamsvcOptionGroup{
					Name: swag.String("test_name"),
				},
				Context: context.TODO(),
			},
		},
		{
			"OptionGroupDelete",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_group/option-group-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionGroupDeleteParams{
				ID:      "option-group-delete-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionGroupList",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_group"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionGroupListParams{
				Context: context.TODO(),
			},
		},
		{
			"OptionGroupRead",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_group/option-group-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&OptionGroupReadParams{
				ID:      "option-group-read-id",
				Context: context.TODO(),
			},
		},
		{
			"OptionGroupUpdate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dhcp/option_group/option-group-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			&OptionGroupUpdateParams{
				ID: "option-group-update-id",
				Body: &models.IpamsvcOptionGroup{
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
			c := initOptionGroupTestClient(s.URL)

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

func initOptionGroupTestClient(url string) ClientService {
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
