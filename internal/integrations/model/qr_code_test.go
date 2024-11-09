package model_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/integrations/model"
)

func TestQRCode(t *testing.T) {
	t.Parallel()

	t.Run("got success when using get json body", func(t *testing.T) {
		t.Parallel()

		input := model.QRCodeInput{
			ExpirationDate:    "ExpirationDate",
			ExternalReference: "ExternalReference",
			Items: []model.Item{
				{
					Description: "Description",
					SkuNumber:   "SKU_NUMBER",
				},
			},
		}

		buffer, err := input.GetJSONBody()

		assert.NoError(t, err)

		var inputToCompare model.QRCodeInput

		err = json.Unmarshal(buffer.Bytes(), &inputToCompare)

		assert.NoError(t, err)

		assert.Equal(t, "ExpirationDate", inputToCompare.ExpirationDate)
		assert.Equal(t, "ExternalReference", inputToCompare.ExternalReference)
		assert.Equal(t, "Description", inputToCompare.Items[0].Description)
		assert.Equal(t, "SKU_NUMBER", inputToCompare.Items[0].SkuNumber)
	})
}
