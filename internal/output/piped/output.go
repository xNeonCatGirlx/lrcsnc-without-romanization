package piped

import (
	"log"
	"lrcsnc/internal/pkg/types"
	"os"
	"strings"
	"time"

	mprislib "github.com/Endg4meZer0/go-mpris"
)

// TODO: JSON output for Waybar and stuff

var outputDestination *os.File = os.Stdout
var outputDestChanged = false
var overwrite = ""
var instrTimer = time.NewTimer(5 * time.Minute)
var writeInstrumental bool = false

func Init() {
	go func() {
		i := 1
		instrTimer.Reset(time.Duration(config.Instrumental.Interval*1000) * time.Millisecond)
		for {
			<-instrTimer.C
			note := config.Instrumental.Symbol
			j := int(config.Instrumental.MaxSymbols + 1)

			// Only update instrumental stuff if the song is playing
			if player.PlaybackStatus == mprislib.PlaybackPlaying && writeInstrumental {
				var stringToPrint string

				switch player.Song.LyricsData.LyricsType {
				case types.LyricsStatePlain:
					stringToPrint = getOutString(config.NoSyncedLyrics)
				case types.LyricsStateNotFound:
					stringToPrint = getOutString(config.SongNotFound)
				case types.LyricsStateInProgress:
					stringToPrint = getOutString(config.GettingLyrics)
				default:
					stringToPrint = getOutString(config.ErrorMessage)
				}

				stringToPrint += " " + strings.Repeat(note, i%j)

				outputDestination.WriteString(stringToPrint + "\n")

				i++
				if i >= j {
					i = 1
				}
			}
			instrTimer.Reset(time.Duration(config.Instrumental.Interval*1000) * time.Millisecond)
		}
	}()
}

func Print(lyric string) {
	if outputDestChanged {
		outputDestination.Truncate(0)
		outputDestination.Seek(0, 0)
	}

	if overwrite != "" {
		return
	}

	if lyric == "" {
		writeInstrumental = true
		instrTimer.Reset(1)
	} else {
		writeInstrumental = false
		instrTimer.Stop()
		outputDestination.WriteString(lyric + "\n")
	}
}

func UpdateDestination(path string) {
	newDest, err := os.Create(path)
	if err != nil {
		log.Println("The output file was set, but I can't create/open it! Permissions issue or wrong path?")
	} else {
		outputDestination = newDest
		outputDestChanged = true
	}
}

func Close() {
	outputDestination.Close()
}

func Overwrite(over string) {
	overwrite = over
	if outputDestChanged {
		outputDestination.Truncate(0)
		outputDestination.Seek(0, 0)
	}
	outputDestination.WriteString(overwrite + "\n")
	go func() {
		<-time.NewTimer(5 * time.Second).C
		overwrite = ""
	}()
}

func Write(s string) {
	
}