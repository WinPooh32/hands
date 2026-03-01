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
func TestTr(t *testing.T) {
	err := i18n.Load()
	require.NoError(t, err)

	tests := []struct {
		name string
		lang string
		key  i18n.Key
		want string
	}{
		{
			name: "English read description",
			lang: "en",
			key:  i18n.ReadDescription,
			want: "Read file contents",
		},
		{
			name: "Russian read description",
			lang: "ru",
			key:  i18n.ReadDescription,
			want: "Чтение содержимого файла",
		},
		{
			name: "English write description",
			lang: "en",
			key:  i18n.WriteDescription,
			want: "Write content to a file",
		},
		{
			name: "English edit description",
			lang: "en",
			key:  i18n.EditDescription,
			want: "Edit a file using search and replace",
		},
		{
			name: "English glob description",
			lang: "en",
			key:  i18n.GlobDescription,
			want: "Find files by pattern",
		},
		{
			name: "English grep description",
			lang: "en",
			key:  i18n.GrepDescription,
			want: "Search for patterns in files",
		},
		{
			name: "English bash description",
			lang: "en",
			key:  i18n.BashDescription,
			want: "Execute a bash command",
		},
		{
			name: "Unknown language (fallbacks to English)",
			lang: "unknown",
			key:  i18n.ReadDescription,
			want: "Read file contents",
		},
		{
			name: "Unknown key",
			lang: "en",
			key:  i18n.Key("nonExistentKey"),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i18n.Language = tt.lang

			got := i18n.Tr(tt.key)
			assert.Equal(t, tt.want, got, "should match expected value")
		})
	}
}

//nolint:paralleltest
func TestTrDefaultLanguage(t *testing.T) {
	err := i18n.Load()
	require.NoError(t, err)

	got := i18n.Tr(i18n.ReadDescription)
	assert.Equal(t, "Read file contents", got, "default should be English")
}
