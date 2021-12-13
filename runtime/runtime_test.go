package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TrimIDPrefix(t *testing.T) {
	testCases := []struct {
		expectedId  string
		pathPattern string
		id          string
	}{
		{
			expectedId:  "test-id",
			pathPattern: "/ipam/ip_space/{id}",
			id:          "ipam/ip_space/test-id",
		},
		{
			expectedId:  "id-beginning-with-slash",
			pathPattern: "/ipam/ip_space/{id}",
			id:          "/ipam/ip_space/id-beginning-with-slash",
		},
		{
			expectedId:  "id-for-copy-path",
			pathPattern: "/ipam/ip_space/{id}/copy",
			id:          "/ipam/ip_space/id-for-copy-path",
		},
		{
			expectedId:  "id-without-prefix",
			pathPattern: "/ipam/ip_space/{id}",
			id:          "id-without-prefix",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expectedId, func(t *testing.T) {
			assert.Equal(t, tc.expectedId, TrimIDPrefix(tc.pathPattern, tc.id))
		})
	}
}
