package auth_zone

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
			"AuthZoneCopy",
			&AuthZoneCopyParams{
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_zone/copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthZoneCreate",
			&AuthZoneCreateParams{
				Body: &models.ConfigAuthZone{
					Fqdn:        swag.String("test.com."),
					PrimaryType: swag.String("cloud"),
					View:        "dns-view-id",
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_zone"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"fqdn\":\"test.com.\",\"primary_type\":\"cloud\",\"view\":\"dns-view-id\"}\n")),
			},
		},
		{
			"AuthZoneDelete",
			&AuthZoneDeleteParams{
				ID:      "auth-zone-delete-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_zone/auth-zone-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthZoneList",
			&AuthZoneListParams{
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
					Path:     "/api/ddi/v1/dns/auth_zone",
					RawQuery: "_fields=field&_filter=filter&_limit=20&_offset=20&_order_by=desc&_page_token=token&_tfilter=tfilter&_torder_by=desc",
				},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthZoneRead",
			&AuthZoneReadParams{
				ID:      "auth-zone-read-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_zone/auth-zone-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"AuthZoneUpdate",
			&AuthZoneUpdateParams{
				ID: "auth-zone-update-id",
				Body: &models.ConfigAuthZone{
					Fqdn:        swag.String("test.com."),
					PrimaryType: swag.String("cloud"),
					View:        "dns-view-id",
					Comment:     "Updated comment",
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/auth_zone/auth-zone-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\",\"fqdn\":\"test.com.\",\"primary_type\":\"cloud\",\"view\":\"dns-view-id\"}\n")),
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
			c := initAuthZoneTestClient(s.URL)

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

func initAuthZoneTestClient(url string) ClientService {
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
