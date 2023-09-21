/*
xendit-invoice-service_test

Testing InvoiceApiService

*/

// Code generated by OpenAPI Generator

package xendit

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	xendit "github.com/xendit/xendit-go/v3"
)

func Test_xendit_InvoiceApiService(t *testing.T) {

	apiKey := os.Getenv("XND_APIKEY")
	if apiKey == "" {
		t.Skip("XND_APIKEY not set")
	}
	
	apiClient := xendit.NewClient(apiKey)

	t.Run("Test InvoiceApiService CreateInvoice", func(t *testing.T) {

		resp, httpRes, err := apiClient.InvoiceApi.CreateInvoice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceApiService ExpireInvoice", func(t *testing.T) {

		var invoiceId string

		resp, httpRes, err := apiClient.InvoiceApi.ExpireInvoice(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceApiService GetInvoiceById", func(t *testing.T) {

		var invoiceId string

		resp, httpRes, err := apiClient.InvoiceApi.GetInvoiceById(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceApiService GetInvoices", func(t *testing.T) {

		resp, httpRes, err := apiClient.InvoiceApi.GetInvoices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}