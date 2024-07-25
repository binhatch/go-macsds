/*
Hella Gutmann - macsDS (Data Services) - API collection

Testing RepairTimesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package github.com/binhatch/go-macds

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_github.com/binhatch/go-macds_RepairTimesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepairTimesAPIService RepairTimes1tree", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepairTimesAPI.RepairTimes1tree(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepairTimesAPIService RepairTimes2details", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepairTimesAPI.RepairTimes2details(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
