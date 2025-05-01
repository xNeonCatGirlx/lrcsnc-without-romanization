package lrclib

import (
	"slices"
	"testing"

	lrclib "lrcsnc/internal/lyrics/providers/lrclib"
	"lrcsnc/internal/pkg/errors"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
)

type Response struct {
	StatusCode int
	Body       string
}

// TestGetLyrics tests the ability to get different kinds of
// lyrics from LrcLib.
func TestGetLyrics(t *testing.T) {
	tests := []struct {
		name  string
		song  structs.Song
		ldata structs.LyricsData
	}{
		{
			name: "existing",
			song: structs.Song{Title: "Earthless", Artists: []string{"Night Verses"}, Album: "From the Gallery of Sleep", Duration: 383},
			ldata: structs.LyricsData{
				Lyrics: []structs.Lyric{
					{Time: 344.18, Text: "\"He is the one who gave me the horse"},
					{Time: 346.74, Text: "So I could ride into the desert and see"},
					{Time: 350.77, Text: "The future.\""},
					{Time: 351.40999999999997, Text: ""},
					{Time: 358.75, Text: "\"He is the one who gave me the horse"},
					{Time: 361.39, Text: "So I could ride into the desert and see"},
					{Time: 365.29, Text: "The future.\""},
					{Time: 366.07, Text: ""},
				},
				LyricsState: types.LyricsStateSynced,
			},
		},
		{
			name: "existing-unicode",
			song: structs.Song{Title: "狼之主", Artists: []string{"塞壬唱片-MSR"}, Album: "敘拉古人OST", Duration: 215},
			ldata: structs.LyricsData{
				Lyrics: []structs.Lyric{
					{Time: 20.48, Text: "You're tough, but it's never been about you"},
					{Time: 23.86, Text: "You're free, but cement your feet, a statue"},
					{Time: 27.02, Text: "Your rules, but you'd rather make up something"},
					{Time: 30.37, Text: "You're dead, you were never good for nothing"},
					{Time: 33.62, Text: "Double negatives leading me in circles"},
					{Time: 36.86, Text: "Twist infinity"},
					{Time: 38.7, Text: "You drive me insane"},
					{Time: 40.78, Text: "Hit it hard, a broken wall"},
					{Time: 44.13, Text: "Hit hard, I gave it all"},
					{Time: 47.36, Text: "Hit hard, a family tie, oh"},
					{Time: 51.46, Text: "But you'd rather just fight"},
					{Time: 53.41, Text: ""},
					{Time: 56.62, Text: ""},
					{Time: 66.11, Text: "Your dirt never washed off in an April shower"},
					{Time: 69.49, Text: "You're crushed by the weight of those you can't devour"},
					{Time: 72.73, Text: "You're armed, but the plan never executed"},
					{Time: 75.82, Text: "You're shocked to let you hold a gun and never shoot it"},
					{Time: 79.17, Text: "Regulating when the rules are simply saturated"},
					{Time: 82.52, Text: "Is it everything, or is it just that I'm insane?"},
					{Time: 86.43, Text: "Hit it hard, a broken wall"},
					{Time: 89.84, Text: "Hit hard, I gave it all"},
					{Time: 93.03, Text: "You tried and failed, all it fell, whoa"},
					{Time: 100.18, Text: ""},
					{Time: 167.69, Text: "All of this to say you lost it all to gain some power"},
					{Time: 170.89, Text: "All of this, just say you plant a seed and kill the flower"},
					{Time: 174.09, Text: "All of this to say you talk your way into your silence"},
					{Time: 177.26, Text: "All of this is just a ploy to force your hand to violence"},
					{Time: 180.36, Text: "It's a waste of time thinking you got tough"},
					{Time: 183.53, Text: "When it's never really been enough"},
					{Time: 185.24, Text: "Am I insane?"},
					{Time: 187.34, Text: "Hit it hard, a broken wall"},
					{Time: 190.52, Text: "Hit hard, I gave it all"},
					{Time: 193.77, Text: "You tried and failed, all it fell, whoa"},
					{Time: 200.68, Text: ""},
				},
				LyricsState: types.LyricsStateSynced,
			},
		},
		{
			name: "existing-plain",
			song: structs.Song{Title: "machine", Artists: []string{"fromjoy"}, Album: "fromjoy", Duration: 103},
			ldata: structs.LyricsData{
				Lyrics: []structs.Lyric{
					{Time: 0, Text: "Filed in place, matching in pace, cowards with batons hiding their face"},
					{Time: 0, Text: "Dispersing counter action, the dividing distraction"},
					{Time: 0, Text: "I watched them get away without a trace"},
					{Time: 0, Text: ""},
					{Time: 0, Text: "Led astray and left behind, your ideals made you blind"},
					{Time: 0, Text: ""},
					{Time: 0, Text: "Stand behind, hold the line, backstab, in the name of worthless pride"},
					{Time: 0, Text: "As they're worshiped by those they lobotomized"},
					{Time: 0, Text: ""},
					{Time: 0, Text: "Insipid and hollow, you're no leader, you just follow"},
					{Time: 0, Text: "The damage you've done, this pain you've caused being swept below"},
					{Time: 0, Text: "There's no indifference"},
					{Time: 0, Text: "I can't ignore this"},
					{Time: 0, Text: "Protect and serve?"},
					{Time: 0, Text: "You're dismissed"},
					{Time: 0, Text: ""},
					{Time: 0, Text: "Black bruises blooming, red and blue glooming"},
					{Time: 0, Text: "The sirens flood as freedom's drained away"},
					{Time: 0, Text: "No redemption as my faith begins to fade"},
				},
				LyricsState: types.LyricsStatePlain,
			},
		},
		{
			name: "existing-instrumental",
			song: structs.Song{Title: "Vice Wave", Artists: []string{"Night Verses"}, Album: "From the Gallery of Sleep", Duration: 300},
			ldata: structs.LyricsData{
				Lyrics:      []structs.Lyric{},
				LyricsState: types.LyricsStateInstrumental,
			},
		},
		{
			name: "non-existing",
			song: structs.Song{Title: "Moonmore", Artists: []string{"Day Choruses"}, Album: "From the Gallery of Minecraft Pictures idk", Duration: 283},
			ldata: structs.LyricsData{
				Lyrics:      []structs.Lyric{},
				LyricsState: types.LyricsStateNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lrclib.Provider{}.Get(tt.song)
			if err != nil && !(tt.ldata.LyricsState == types.LyricsStateNotFound && err == errors.ErrLyricsNotFound) {
				t.Errorf("[tests/lyrics/providers/lrclib/get/%v] Error: %v", tt.name, err)
				return
			}
			if !slices.Equal(got.Lyrics, tt.ldata.Lyrics) || got.LyricsState != tt.ldata.LyricsState {
				t.Errorf("[tests/lyrics/providers/lrclib/get/%v] Received %v, want %v", tt.name, got, tt.ldata)
			}
		})
	}
}
