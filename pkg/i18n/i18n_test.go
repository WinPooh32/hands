package i18n_test

import (
	"testing"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	err := i18n.Load()
	require.NoError(t, err)
}

//nolint:paralleltest
func TestTrDefaultLanguage(t *testing.T) {
	err := i18n.Load()
	require.NoError(t, err)

	gotDefault := i18n.Tr(i18n.ReadDescription)

	i18n.Language = "ru"
	gotRu := i18n.Tr(i18n.ReadDescription)

	assert.NotEqual(t, gotDefault, gotRu, "default should be English")
}
