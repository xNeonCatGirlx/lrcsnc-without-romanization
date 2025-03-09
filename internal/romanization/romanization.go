package romanization

import (
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"strings"
	"unicode"

	zh "github.com/mozillazg/go-pinyin"
	jp "github.com/sarumaj/go-kakasi"
	kr "github.com/srevinsaju/korean-romanizer-go"
)

// TODO: there may be songs with mixed languages, so we may need to romanize each line separately

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
func RomanizeLyrics(lyrics []structs.Lyric) {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	if !global.Config.C.Lyrics.Romanization.IsEnabled() {
		return
	}

	lang := getLang(lyrics)
	if lang == 0 {
		return
	}

	for i := range lyrics {
		var rstr string = ""
		if len(lyrics[i].Text) != 0 {
			rstr = romanize(lyrics[i].Text, lang)
		}
		lyrics[i].Text = rstr
	}
}

// Returns the first found language supported by romanization module,
// or falls back to LanguageDefault if no supported language is found
func getLang(lyrics []structs.Lyric) Language {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	if global.Config.C.Lyrics.Romanization.Japanese {
		for i := range lyrics {
			if isChar(lyrics[i].Text, jpUnicodeRangeTable) {
				return LanguageJapanese
			}
		}
	}
	if global.Config.C.Lyrics.Romanization.Korean {
		for i := range lyrics {
			if isChar(lyrics[i].Text, krUnicodeRangeTable) {
				return LanguageKorean
			}
		}
	}
	if global.Config.C.Lyrics.Romanization.Chinese {
		for i := range lyrics {
			if isChar(lyrics[i].Text, zhUnicodeRangeTable) {
				return LanguageChinese
			}
		}
	}
	return LanguageDefault
}

// Returns a romanized string based on the provided language
func romanize(str string, lang Language) (out string) {
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
