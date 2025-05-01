package langs

import (
	"lrcsnc/internal/pkg/log"
	"os/exec"
	"strings"
)

type RomanizerJp struct{}

func (RomanizerJp) Romanize(str string) string {
	if str == "" {
		return ""
	}

	cmd := exec.Command("kakasi", "-i", "utf8", "-o", "utf8", "-Ha", "-Ka", "-Ja", "-Ea", "-ka", "-s")
	cmd.Stdin = strings.NewReader(str)
	out, err := cmd.Output()
	if err != nil {
		log.Error("romanization/langs/jp", "Error executing kakasi command: "+err.Error())
		return str
	}
	return strings.TrimSpace(string(out))
}
