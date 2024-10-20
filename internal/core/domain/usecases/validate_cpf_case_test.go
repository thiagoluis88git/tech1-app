package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCPFUseCase(t *testing.T) {
	t.Parallel()

	t.Run("got true when validating right CPF", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateCPFUseCase()

		cleanedCPF, validated := sut.Execute("171.079.720-73")

		assert.True(t, validated)
		assert.Equal(t, "17107972073", cleanedCPF)
	})

	t.Run("got true when validating right cleaned CPF", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateCPFUseCase()

		cleanedCPF, validated := sut.Execute("17107972073")

		assert.True(t, validated)
		assert.Equal(t, "17107972073", cleanedCPF)
	})

	t.Run("got false when validating wrong CPF", func(t *testing.T) {
		t.Parallel()

		sut := NewValidateCPFUseCase()

		cleanedCPF, validated := sut.Execute("172.079.720-73")

		assert.False(t, validated)
		assert.Equal(t, "17207972073", cleanedCPF)
	})
}
