package range_operations

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
			"RangeCreate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"end\":\"192.168.1.30\",\"space\":\"ip-space-id\",\"start\":\"192.168.1.15\"}\n")),
			},
			&RangeCreateParams{
				Body: &models.IpamsvcRange{
					Start: swag.String("192.168.1.15"),
					End:   swag.String("192.168.1.30"),
					Space: swag.String("ip-space-id"),
				},
				Context: context.TODO(),
			},
		},
		{
			"RangeCreateNextAvailableIP",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range/range-create-next-available-ip-id/nextavailableip"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeCreateNextAvailableIPParams{
				ID:      "range-create-next-available-ip-id",
				Context: context.TODO(),
			},
		},
		{
			"RangeDelete",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range/range-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeDeleteParams{
				ID:      "range-delete-id",
				Context: context.TODO(),
			},
		},
		{
			"RangeList",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeListParams{
				Context: context.TODO(),
			},
		},
		{
			"RangeListNextAvailableIP",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range/range-list-next-available-ip-id/nextavailableip"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeListNextAvailableIPParams{
				ID:      "range-list-next-available-ip-id",
				Context: context.TODO(),
			},
		},
		{
			"RangeRead",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range/range-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeReadParams{
				ID:      "range-read-id",
				Context: context.TODO(),
			},
		},
		{
			"RangeUpdate",
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/range/range-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			&RangeUpdateParams{
				ID:      "range-update-id",
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
			c := initRangeOperationsTestClient(s.URL)

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

func initRangeOperationsTestClient(url string) ClientService {
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
