package osuapi

import "strconv"

// Languages
const (
	LanguageAny Language = iota
	LanguageOther
	LanguageEnglish
	LanguageJapanese
	LanguageChinese
	LanguageInstrumental
	LanguageKorean
	LanguageFrench
	LanguageGerman
	LanguageSwedish
	LanguageSpanish
	LanguageItalian
)

// Language is the language of a beatmap's song.
type Language int

var languageString = [...]string{
	"any",
	"other",
	"English",
	"Japanese",
	"Chinese",
	"instrumental",
	"Korean",
	"French",
	"German",
	"Swedish",
	"Spanish",
	"Italian",
}

func (l Language) String() string {
	if l >= 0 && int(l) < len(languageString) {
		return languageString[l]
	}
	return strconv.Itoa(int(l))
}
