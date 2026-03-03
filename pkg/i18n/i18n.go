package i18n

import (
	"embed"
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
)

var Language string

type Key string

const (
	ReadDescription   Key = "readDescription"
	ReadArgPath       Key = "readArgPath"
	ReadArgEncoding   Key = "readArgEncoding"
	WriteDescription  Key = "writeDescription"
	WriteArgPath      Key = "writeArgPath"
	WriteArgContent   Key = "writeArgContent"
	WriteArgEncoding  Key = "writeArgEncoding"
	EditDescription   Key = "editDescription"
	EditArgPath       Key = "editArgPath"
	EditArgSearch     Key = "editArgSearch"
	EditArgReplace    Key = "editArgReplace"
	GlobDescription   Key = "globDescription"
	GlobArgPattern    Key = "globArgPattern"
	GlobArgDir        Key = "globArgDir"
	GrepDescription   Key = "grepDescription"
	GrepArgPattern    Key = "grepArgPattern"
	GrepArgPath       Key = "grepArgPath"
	GrepArgIgnoreCase Key = "grepArgIgnoreCase"
	BashDescription   Key = "bashDescription"
	BashArgCommand    Key = "bashArgCommand"
	BashArgWorkingDir Key = "bashArgWorkingDir"
)

//go:embed locales/*.yaml
var localesFS embed.FS

var locales map[string]locale

var mut sync.Mutex

type locale struct {
	Name   string
	Values map[Key]string
}

func Load() error {
	mut.Lock()
	defer mut.Unlock()

	locales = make(map[string]locale)

	entries, err := localesFS.ReadDir("locales")
	if err != nil {
		return fmt.Errorf("failed to read locales directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		data, err := localesFS.ReadFile(fmt.Sprintf("locales/%s", entry.Name()))
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", entry.Name(), err)
		}

		var l locale
		if err := yaml.Unmarshal(data, &l); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", entry.Name(), err)
		}

		// Extract language code from filename (e.g., en.yaml -> en)
		lang := entry.Name()
		if len(lang) >= 5 && lang[len(lang)-5:] == ".yaml" {
			lang = lang[:len(lang)-5]
		}

		l.Name = lang
		locales[lang] = l
	}

	return nil
}

func Tr(key Key) string {
	mut.Lock()
	defer mut.Unlock()

	if Language == "" {
		Language = "en"
	}

	loc, ok := locales[Language]
	if !ok {
		// Fallback to English
		loc, ok = locales["en"]
		if !ok {
			return ""
		}
	}

	value, ok := loc.Values[key]
	if !ok && Language != "en" {
		// Fallback to English if not found in current language
		enLoc, ok := locales["en"]
		if !ok {
			return ""
		}

		value = enLoc.Values[key]
	}

	return value
}
