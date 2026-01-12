package middlewares

import (
	"net/http"
	"strings"
)

const (
	LangEn = "en"
	LangId = "id"
)

func GetLanguage(r *http.Request) string {
	// Check Query ID
	lang := r.URL.Query().Get("lang")
	if lang == LangEn || lang == LangId {
		return lang
	}

	// Check Header Accept Language
	headerLang := r.Header.Get("Accept-Language")
	if strings.HasPrefix(headerLang, LangId) {
		return LangId
	}

	// Default to English
	return LangEn
}