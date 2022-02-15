package address

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
			"AddressCreate",
			&AddressCreateParams{
				Body: &models.IpamsvcAddress{
					Address: swag.String("192.168.1.15"),
					Space:   swag.String("ip-space-id"),
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"address\":\"192.168.1.15\",\"space\":\"ip-space-id\"}\n")),
			},
		},
		{
			"AddressDelete",
			&AddressDeleteParams{
				ID:      "address-delete-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address/address-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressList",
			&AddressListParams{
				Fields:       swag.String("field"),
				Filter:       swag.String("filter"),
				Limit:        swag.Int64(int64(20)),
				Offset:       swag.Int64(int64(20)),
				OrderBy:      swag.String("desc"),
				PageToken:    swag.String("token"),
				Tfilter:      swag.String("tfilter"),
				TorderBy:     swag.String("desc"),
				AddressState: swag.String("state"),
				Scope:        swag.String("scope"),
				Context:      context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path: "/api/ddi/v1/ipam/address",
					RawQuery: "_fields=field&_filter=filter&_limit=20&_offset=20&_order_by=desc&_page_token=token" +
						"&_tfilter=tfilter&_torder_by=desc&address_state=state&scope=scope",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressRead",
			&AddressReadParams{
				ID:      "address-read-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address/address-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressUpdate",
			&AddressUpdateParams{
				ID: "address-update-id",
				Body: &models.IpamsvcAddress{
					Comment: "New comment",
				},
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address/address-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"New comment\"}\n")),
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
			c := initAddressTestClient(s.URL)

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

func initAddressTestClient(url string) ClientService {
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
