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
		testMethodParams interface{}
		expectedRequest  http.Request
	}{
		{
			"SubnetCopy",
			&SubnetCopyParams{
				ID:      "subnet-copy-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-copy-id/copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetCreate",
			&SubnetCreateParams{
				Body: &models.IpamsvcSubnet{
					Address: swag.String("192.168.1.0"),
					Cidr:    24,
					Space:   swag.String("ip-space-id"),
				},
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"address\":\"192.168.1.0\",\"cidr\":24,\"space\":\"ip-space-id\"}\n")),
			},
		},
		{
			"SubnetCreateNextAvailableIP",
			&SubnetCreateNextAvailableIPParams{
				ID:      "subnet-create-next-available-ip-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-create-next-available-ip-id/nextavailableip"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetDelete",
			&SubnetDeleteParams{
				ID:      "subnet-delete-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetList",
			&SubnetListParams{
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
					Path:     "/api/ddi/v1/ipam/subnet",
					RawQuery: "_fields=field&_filter=filter&_limit=20&_offset=20&_order_by=desc&_page_token=token&_tfilter=tfilter&_torder_by=desc",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetListNextAvailableIP",
			&SubnetListNextAvailableIPParams{
				ID:         "subnet-list-next-available-ip-id",
				Contiguous: swag.Bool(true),
				Count:      swag.Int32(5),
				Context:    context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/subnet/subnet-list-next-available-ip-id/nextavailableip",
					RawQuery: "contiguous=true&count=5",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetRead",
			&SubnetReadParams{
				ID:      "subnet-read-id",
				Fields:  swag.String("field"),
				Context: context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/subnet/subnet-read-id",
					RawQuery: "_fields=field",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"SubnetUpdate",
			&SubnetUpdateParams{
				ID: "subnet-update-id",
				Body: &models.IpamsvcSubnet{
					Comment: "Updated comment",
				},
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/subnet/subnet-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
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

			// Compose test method call parameters
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
