package romanization_test

import (
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/romanization"
	"os/exec"
	"testing"
)

func TestRomanize(t *testing.T) {
	global.Config.C.Lyrics.Romanization.Japanese = true
	global.Config.C.Lyrics.Romanization.Chinese = true
	global.Config.C.Lyrics.Romanization.Korean = true

	ogJpLyric := "ああ？私に近づいてるの？"

	jpLyrics := []structs.Lyric{{Text: ogJpLyric}}
	krLyrics := []structs.Lyric{{Text: "어? 나한테 다가오니?"}}
	zhLyrics := []structs.Lyric{{Text: "哦？你在接近我吗？"}}
	romanLyrics := []structs.Lyric{{Text: "france?!?"}}
	romanization.Romanize(jpLyrics)
	romanization.Romanize(krLyrics)
	romanization.Romanize(zhLyrics)
	romanization.Romanize(romanLyrics)

	rightAnswerJapanese := []structs.Lyric{{Text: "Aa? Watashi ni chikazu iteruno?"}}
	rightAnswerKorean := []structs.Lyric{{Text: "Eo? Nahante dagaoni?"}}
	rightAnswerChinese := []structs.Lyric{{Text: "Ó? Nǐzàijiējìnwǒma?"}}
	rightAnswerDefault := []structs.Lyric{{Text: "france?!?"}}

	if _, err := exec.LookPath("kakasi"); (err == nil && jpLyrics[0] != rightAnswerJapanese[0]) ||
		(err != nil && jpLyrics[0].Text != ogJpLyric) ||
		krLyrics[0] != rightAnswerKorean[0] ||
		zhLyrics[0] != rightAnswerChinese[0] ||
		romanLyrics[0] != rightAnswerDefault[0] {
		t.Errorf(
			"[tests/romanization/TestRomanize] ERROR: Expected \"%v\", \"%v\", \"%v\" and \"%v\"; received \"%v\", \"%v\", \"%v\" and \"%v\"",
			rightAnswerJapanese[0], rightAnswerKorean[0], rightAnswerChinese[0], rightAnswerDefault[0],
			jpLyrics[0], krLyrics[0], zhLyrics[0], romanLyrics[0])
	}
}
