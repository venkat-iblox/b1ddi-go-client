package record

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
		testMethodName   string
		testMethodParams interface{}
		expectedRequest  http.Request
	}{
		{
			"RecordCreate",
			&RecordCreateParams{
				Body: &models.DataRecord{
					Zone:       "dns-zone-id",
					NameInZone: "record",
					Rdata:      map[string]string{"address": "192.168.1.15"},
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name_in_zone\":\"record\",\"rdata\":{\"address\":\"192.168.1.15\"},\"zone\":\"dns-zone-id\"}\n")),
			},
		},
		{
			"RecordDelete",
			&RecordDeleteParams{
				ID:      "record-delete-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record/record-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"RecordList",
			&RecordListParams{
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"RecordRead",
			&RecordReadParams{
				ID:      "record-read-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record/record-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"RecordSOASerialIncrement",
			&RecordSOASerialIncrementParams{
				ID:      "record-soa-serial-increment-id",
				Body:    &models.DataSOASerialIncrementRequest{},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record/record-soa-serial-increment-id/serial_increment"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{}\n")),
			},
		},
		{
			"RecordUpdate",
			&RecordUpdateParams{
				ID: "record-update-id",
				Body: &models.DataRecord{
					Zone:       "dns-zone-id",
					NameInZone: "record",
					Rdata:      map[string]string{"address": "192.168.1.15"},
					Comment:    "Updated comment",
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/record/record-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\",\"name_in_zone\":\"record\",\"rdata\":{\"address\":\"192.168.1.15\"},\"zone\":\"dns-zone-id\"}\n")),
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
			c := initRecordTestClient(s.URL)

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

func initRecordTestClient(url string) ClientService {
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
