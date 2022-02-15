package address_block

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
			"AddressBlockCopy",
			&AddressBlockCopyParams{
				ID:      "address-block-copy-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-copy-id/copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockCreate",
			&AddressBlockCreateParams{
				Body: &models.IpamsvcAddressBlock{
					Address: swag.String("192.168.1.0"),
					Cidr:    24,
				},
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"address\":\"192.168.1.0\",\"cidr\":24}\n")),
			},
		},
		{
			"AddressBlockCreateNextAvailableAB",
			&AddressBlockCreateNextAvailableABParams{
				ID:      "address-block-next-available-address-block-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-next-available-address-block-id/nextavailableaddressblock"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockCreateNextAvailableIP",
			&AddressBlockCreateNextAvailableIPParams{
				ID:      "address-block-next-available-ip-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-next-available-ip-id/nextavailableip"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockCreateNextAvailableSubnet",
			&AddressBlockCreateNextAvailableSubnetParams{
				ID:      "address-block-next-available-subnet-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-next-available-subnet-id/nextavailablesubnet"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockDelete",
			&AddressBlockDeleteParams{
				ID:      "address-block-delete-id",
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockList",
			&AddressBlockListParams{
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
					Path:     "/api/ddi/v1/ipam/address_block",
					RawQuery: "_fields=field&_filter=filter&_limit=20&_offset=20&_order_by=desc&_page_token=token&_tfilter=tfilter&_torder_by=desc",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockListNextAvailableAB",
			&AddressBlockListNextAvailableABParams{
				Cidr:    swag.Int32(25),
				Comment: swag.String("test-comment"),
				Count:   swag.Int32(5),
				Name:    swag.String("test-name"),
				ID:      "address-block-next-available-address-block-id",
				Context: context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/address_block/address-block-next-available-address-block-id/nextavailableaddressblock",
					RawQuery: "cidr=25&comment=test-comment&count=5&name=test-name",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockListNextAvailableIP",
			&AddressBlockListNextAvailableIPParams{
				Contiguous: swag.Bool(true),
				Count:      swag.Int32(5),
				ID:         "address-block-next-available-ip-id",
				Context:    context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/address_block/address-block-next-available-ip-id/nextavailableip",
					RawQuery: "contiguous=true&count=5",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockListNextAvailableSubnet",
			&AddressBlockListNextAvailableSubnetParams{
				Cidr:     swag.Int32(25),
				Comment:  swag.String("comment"),
				Count:    swag.Int32(2),
				DhcpHost: swag.String("dhcp-host"),
				Name:     swag.String("name"),
				ID:       "address-block-next-available-subnet-id",
				Context:  context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/address_block/address-block-next-available-subnet-id/nextavailablesubnet",
					RawQuery: "cidr=25&comment=comment&count=2&dhcp_host=dhcp-host&name=name",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockRead",
			&AddressBlockReadParams{
				Fields:  swag.String("field"),
				ID:      "address-block-read-id",
				Context: context.TODO(),
			},
			http.Request{
				URL: &url.URL{
					Path:     "/api/ddi/v1/ipam/address_block/address-block-read-id",
					RawQuery: "_fields=field",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AddressBlockUpdate",
			&AddressBlockUpdateParams{
				ID:      "address-block-update-id",
				Body:    &models.IpamsvcAddressBlock{Comment: "Updated comment"},
				Context: context.TODO(),
			},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/address_block/address-block-update-id"},
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
			c := initAddressBlockTestClient(s.URL)

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

func initAddressBlockTestClient(url string) ClientService {
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
