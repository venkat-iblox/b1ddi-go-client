package subnet

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
			"SubnetCopy",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-copy-id/copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetCopyParams{
				ID:      "subnet-copy-id",
				Context: context.TODO(),
			},
		},
		{
			"SubnetCreate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"address\":\"192.168.1.0\",\"cidr\":24,\"space\":\"ip-space-id\"}\n")),
			},
			&SubnetCreateParams{
				Body: &models.IpamsvcSubnet{
					Address: swag.String("192.168.1.0"),
					Cidr:    24,
					Space:   swag.String("ip-space-id"),
				},
				Context: context.TODO(),
			},
		},
		{
			"SubnetCreateNextAvailableIP",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-create-next-available-ip-id/nextavailableip"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetCreateNextAvailableIPParams{
				ID:      "subnet-create-next-available-ip-id",
				Context: context.TODO(),
			},
		},
		{
			"SubnetDelete",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetDeleteParams{
				ID:      "subnet-delete-id",
				Context: context.TODO(),
			},
		},
		{
			"SubnetList",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetListParams{
				Context: context.TODO(),
			},
		},
		{
			"SubnetListNextAvailableIP",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-list-next-available-ip-id/nextavailableip"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetListNextAvailableIPParams{
				ID:      "subnet-list-next-available-ip-id",
				Context: context.TODO(),
			},
		},
		{
			"SubnetRead",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&SubnetReadParams{
				ID:      "subnet-read-id",
				Context: context.TODO(),
			},
		},
		{
			"SubnetUpdate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			&SubnetUpdateParams{
				ID: "subnet-update-id",
				Body: &models.IpamsvcSubnet{
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
			c := initSubnetTestClient(s.URL)

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

func initSubnetTestClient(url string) ClientService {
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
