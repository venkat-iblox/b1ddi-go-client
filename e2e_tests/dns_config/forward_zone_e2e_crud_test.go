//go:build e2e
// +build e2e

package e2e_tests

import (
	"context"
	"testing"
	"time"

	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_zone"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForwardZoneCRUD(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	testCases := map[string]struct {
		instance   interface{}
		updateFunc interface{}
	}{}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			createResp, err := cli.ForwardZone.ForwardZoneCreate(
				&forward_zone.ForwardZoneCreateParams{Body: tc.instance, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, createResp.Payload.Result.ID)

			readResp, err := cli.ForwardZone.ForwardZoneRead(
				&forward_zone.ForwardZoneReadParams{ID: createResp.Payload.Result.ID, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.Equal(t, createResp.Payload.Result, readResp.Payload.Result)

			tc.updateFunc(tc.instance)
			updateResp, err := cli.ForwardZone.ForwardZoneUpdate(
				&forward_zone.ForwardZoneUpdateParams{ID: createResp.Payload.Result.ID, Body: tc.instance, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, updateResp)

			deleteResp, err := cli.ForwardZone.ForwardZoneDelete(
				&forward_zone.ForwardZoneDeleteParams{ID: updateResp.Payload.Result.ID, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, deleteResp)
		})
	}
}
