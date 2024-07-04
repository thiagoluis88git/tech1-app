package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

func TestValidateProductCategoryUseCase(t *testing.T) {
	t.Run("got true when validating product with right parameters", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateProductCategoryUseCase()

		validated := sut.Execute(domain.ProductForm{
			Category: "Lanche",
		})

		assert.True(t, validated)
	})

	t.Run("got true when validating combo with right parameters", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateProductCategoryUseCase()

		validated := sut.Execute(domain.ProductForm{
			Category:         "Combo",
			ComboProductsIds: &[]uint{1, 2},
		})

		assert.True(t, validated)
	})

	t.Run("got false when validating combo with wrong parameters", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateProductCategoryUseCase()

		validated := sut.Execute(domain.ProductForm{
			Category: "Combo",
		})

		assert.False(t, validated)
	})
}
