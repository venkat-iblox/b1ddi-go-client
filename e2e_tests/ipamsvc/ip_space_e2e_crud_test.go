//go:build e2e
// +build e2e

package e2e_tests

import (
	"context"
	"github.com/go-openapi/swag"
	"testing"
	"time"

	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIPSpaceCRUD(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	testCases := map[string]struct {
		instance   *models.IpamsvcIPSpace
		updateFunc func(instance *models.IpamsvcIPSpace)
	}{
		"CreateReadUpdateNameDelete": {
			instance: &models.IpamsvcIPSpace{
				Name: swag.String("ip_space_go_client_e2e_test"),
			},
			updateFunc: func(instance *models.IpamsvcIPSpace) {
				instance.Name = swag.String("ip_space_updated_go_client_e2e_test")
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			createResp, err := cli.IPSpace.IPSpaceCreate(
				&ip_space.IPSpaceCreateParams{Body: tc.instance, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, createResp.Payload.Result.ID)

			readResp, err := cli.IPSpace.IPSpaceRead(
				&ip_space.IPSpaceReadParams{ID: createResp.Payload.Result.ID, Context: ctx}, nil,
			)
			assert.NoError(t, err)
			assert.Equal(t, createResp.Payload.Result, readResp.Payload.Result)

			tc.updateFunc(tc.instance)
			updateResp, err := cli.IPSpace.IPSpaceUpdate(
				&ip_space.IPSpaceUpdateParams{ID: createResp.Payload.Result.ID, Body: tc.instance, Context: ctx}, nil,
			)
			assert.NoError(t, err)
			assert.NotNil(t, updateResp)

			deleteResp, err := cli.IPSpace.IPSpaceDelete(
				&ip_space.IPSpaceDeleteParams{ID: updateResp.Payload.Result.ID, Context: ctx}, nil,
			)
			assert.NoError(t, err)
			assert.NotNil(t, deleteResp)
		})
	}
}
