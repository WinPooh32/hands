package i18n

import (
	"embed"
)

type Key string

const (
	ReadDescription  Key = "readDescription"
	WriteDescription Key = "writeDescription"
	EditDescription  Key = "editDescription"
	GlobDescription  Key = "globDescription"
	GrepDescription  Key = "grepDescription"
	BashDescription  Key = "bashDescription"
)

//go:embed locales/*.yaml
var localesDir embed.FS

var locales map[string]locale

type locale struct {
	Name   string
	Values map[Key]string
}

func Load() error {
	panic("TODO")
}

func Tr(lang string, key Key) string {
	panic("TODO")
}
