package romanization

import (
	"strings"
	"unicode"

	"lrcsnc/internal/pkg/global"

	zh "github.com/mozillazg/go-pinyin"
	jp "github.com/sarumaj/go-kakasi"
	kr "github.com/srevinsaju/korean-romanizer-go"
)

// Supported languages are listed here
type Language uint

var (
	LanguageDefault  Language = 0
	LanguageJapanese Language = 1
	LanguageKorean   Language = 2
	LanguageChinese  Language = 3
)

// Unicode range tables accordingly are listed here
var jpUnicodeRangeTable = []*unicode.RangeTable{
	unicode.Hiragana,
	unicode.Katakana,
	unicode.Diacritic,
}

var krUnicodeRangeTable = []*unicode.RangeTable{
	unicode.Hangul,
}

var zhUnicodeRangeTable = []*unicode.RangeTable{
	unicode.Ideographic,
	unicode.Han,
}

// Returns romanized lyrics (or the same lyrics if the language is not supported)
func RomanizeLyrics(strs []string) []string {
	global.CurrentConfig.Mutex.Lock()
	defer global.CurrentConfig.Mutex.Unlock()

	if !global.CurrentConfig.Config.Lyrics.Romanization.IsEnabled() {
		return strs
	}

	lang := GetLang(strs)
	if lang == 0 {
		return strs
	}

	out := make([]string, 0, len(strs))
	for _, str := range strs {
		var rstr string = ""
		if len(str) != 0 {
			rstr = Romanize(str, lang)
		}
		out = append(out, rstr)
	}
	return out
}

// Returns the first found language supported by romanization module, or falls back to LanguageDefault
func GetLang(lyrics []string) Language {
	global.CurrentConfig.Mutex.Lock()
	defer global.CurrentConfig.Mutex.Unlock()

	if global.CurrentConfig.Config.Lyrics.Romanization.Japanese {
		for _, l := range lyrics {
			if isChar(l, jpUnicodeRangeTable) {
				return LanguageJapanese
			}
		}
	}
	if global.CurrentConfig.Config.Lyrics.Romanization.Korean {
		for _, l := range lyrics {
			if isChar(l, krUnicodeRangeTable) {
				return LanguageKorean
			}
		}
	}
	if global.CurrentConfig.Config.Lyrics.Romanization.Chinese {
		for _, l := range lyrics {
			if isChar(l, zhUnicodeRangeTable) {
				return LanguageChinese
			}
		}
	}
	return LanguageDefault
}

// Returns a romanized string based on the provided language
func Romanize(str string, lang Language) (out string) {
	switch lang {
	case LanguageJapanese:
		kakasiConverter, err := jp.NewKakasi()
		if err != nil {
			// TODO: logger
			panic(err)
		}

		converted, err := kakasiConverter.Convert(str)
		if err != nil {
			// TODO: logger
			panic(err)
		}

		out = converted.Romanize()
		if unicode.IsLower(rune(out[0])) {
			out = strings.ToUpper(out[:1]) + out[1:]
		}
	case LanguageKorean:
		r := kr.NewRomanizer(str)
		out = r.Romanize()
		if unicode.IsLower(rune(out[0])) {
			out = strings.ToUpper(out[:1]) + out[1:]
		}
	case LanguageChinese:
		out = zhCharToPinyin(str)
		if unicode.IsLower(rune(out[0])) {
			out = strings.ToUpper(out[:1]) + out[1:]
		}
	}
	return out
}

func isChar(s string, rangeTable []*unicode.RangeTable) bool {
	for _, r := range s {
		if unicode.IsOneOf(rangeTable, r) {
			return true
		}
	}
	return false
}

func zhCharToPinyin(p string) (s string) {
	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			s += string(zh.Pinyin(string(r), zh.NewArgs())[0][0])
		} else {
			s += string(r)
		}
	}
	return
}
