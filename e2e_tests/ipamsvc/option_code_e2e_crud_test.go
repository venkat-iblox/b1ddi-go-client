//go:build e2e
// +build e2e

package e2e_tests

import (
	"context"
	"testing"
	"time"

	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionCodeCRUD(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	testCases := map[string]struct {
		instance   interface{}
		updateFunc interface{}
	}{}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			createResp, err := cli.OptionCode.OptionCodeCreate(
				&option_code.OptionCodeCreateParams{Body: tc.instance, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, createResp.Payload.Result.ID)

			readResp, err := cli.OptionCode.OptionCodeRead(
				&option_code.OptionCodeReadParams{ID: createResp.Payload.Result.ID, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.Equal(t, createResp.Payload.Result, readResp.Payload.Result)

			tc.updateFunc(tc.instance)
			updateResp, err := cli.OptionCode.OptionCodeUpdate(
				&option_code.OptionCodeUpdateParams{ID: createResp.Payload.Result.ID, Body: tc.instance, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, updateResp)

			deleteResp, err := cli.OptionCode.OptionCodeDelete(
				&option_code.OptionCodeDeleteParams{ID: updateResp.Payload.Result.ID, Context: ctx}, nil,
			)
			require.NoError(t, err)
			assert.NotNil(t, deleteResp)
		})
	}
}
